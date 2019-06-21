package gitlab

import (
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/factorysh/minasan/metrics"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	gitlab "github.com/xanzy/go-gitlab"
)

// Client for Gitlab REST API
type Client struct {
	*gitlab.Client
}

var clientCache = cache.New(5*time.Minute, 10*time.Minute)
var groupMembersCache = cache.New(5*time.Minute, 60*time.Minute)

// NewClientWithGitlabPrivateToken returns a new Client with a Gitlab's private token
func NewClientWithGitlabPrivateToken(client *http.Client, gitlabDomain string, privateToken string) *Client {
	cached, found := clientCache.Get(gitlabDomain)
	if found {
		return &Client{cached.(*gitlab.Client)}
	}
	gl := gitlab.NewClient(client, privateToken)
	gl.SetBaseURL("https://" + gitlabDomain + "/api/v4")
	clientCache.SetDefault(gitlabDomain, gl)
	return &Client{gl}
}

// NewClientFromEnv returns a new Client from environments
func NewClientFromEnv(client *http.Client) *Client {
	return NewClientWithGitlabPrivateToken(client, os.Getenv("GITLAB_DOMAIN"), os.Getenv("GITLAB_PRIVATE_TOKEN"))
}

// MailsFromGroupProject returns distincts mails from a project and its group
func (c *Client) MailsFromGroupProject(group, project string) ([]string, error) {
	const level = 40

	// Get the group members in the cache
	groupMembers, expTime, found := groupMembersCache.GetWithExpiration(group)
	// If the groupMembers was not found in the cache or found but expired
	// Else -> continue with the cache
	if !found || (found && cache.Item{groupMembers, expTime.Unix()}.Expired()) {
		// Works for gitlab 9, but documentation talks about https://docs.gitlab.com/ce/api/members.html#list-all-members-of-a-group-or-project-including-inherited-members
		// It doesn't work with curl + private token, and go-gitlab seems to not implement it
		groupMembers, resp, err := c.Groups.ListGroupMembers(group, &gitlab.ListGroupMembersOptions{})
		// If gitlab is unavailable and the groupMembers was not found in the cache -> error
		// Else if gitlab is available -> put the new groupMembers in the cache
		// Else -> continue with the expired cache
		if err != nil && !found {
			log.WithField("response", resp).WithError(err).Error("MailsFromGroupProject")
			if resp.StatusCode == 404 {
				metrics.WrongProjectCounter.Inc()
			}
			return nil, err
		} else if err == nil {
			groupMembersCache.SetDefault(group, groupMembers)
		}
	}
	mails := make(map[string]interface{})
	for _, member := range groupMembers.([]*gitlab.GroupMember) {
		if member.AccessLevel < level || member.State != "active" {
			continue
		}
		user, resp, err := c.Users.GetUser(member.ID)
		if err != nil {
			log.WithField("response", resp).WithError(err).Error("MailsFromGroup")
			return nil, err
		}
		log.WithField("name", user.Name).WithField("email", user.Email).WithField("group", group).Debug("Users from group")
		if user.Email == "" {
			log.WithField("name", user.Name).Warning("User with an empty email")
		}
		mails[user.Email] = true
	}

	id := strings.Join([]string{group, project}, "/")
	members, resp, err := c.ProjectMembers.ListProjectMembers(id, &gitlab.ListProjectMembersOptions{})
	if err != nil {
		log.WithField("response", resp).WithError(err).Error("ListProjectMembers")
		return nil, err
	}
	for _, member := range members {
		if member.AccessLevel < level || member.State != "active" {
			continue
		}
		user, resp, err := c.Users.GetUser(member.ID)
		if err != nil {
			log.WithField("response", resp).WithError(err).Error("MailsFromProject")
			return nil, err
		}
		log.WithField("name", user.Name).WithField("email", user.Email).WithField("project", project).Debug("Users from project")
		if user.Email == "" {
			log.WithField("name", user.Name).Warning("Member with an empty email")
		}
		mails[user.Email] = true
	}
	smails := make([]string, len(mails))
	i := 0
	for key := range mails {
		smails[i] = key
		i++
	}
	sort.Strings(smails)
	return smails, nil
}

// Ping my name
func (c *Client) Ping() (string, error) {
	user, _, err := c.Users.CurrentUser()
	if err != nil {
		return "", err
	}
	return user.Name, nil
}
