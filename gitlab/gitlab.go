package gitlab

import (
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/factorysh/minasan/cache"
	"github.com/factorysh/minasan/metrics"
	log "github.com/sirupsen/logrus"
	gitlab "github.com/xanzy/go-gitlab"
)

// Client for Gitlab REST API
type Client struct {
	*gitlab.Client
}

// var groupMembersCache = cache.NewCache()

// NewClientWithGitlabPrivateToken returns a new Client with a Gitlab's private token
func NewClientWithGitlabPrivateToken(client *http.Client, gitlabDomain string, privateToken string) *Client {
	gl := gitlab.NewClient(client, privateToken)
	gl.SetBaseURL("https://" + gitlabDomain + "/api/v4")
	return &Client{gl}
}

// NewClientFromEnv returns a new Client from environments
func NewClientFromEnv(client *http.Client) *Client {
	return NewClientWithGitlabPrivateToken(client, os.Getenv("GITLAB_DOMAIN"), os.Getenv("GITLAB_PRIVATE_TOKEN"))
}

func (c *Client) GetGitlabGroupMembers(key string) (interface{}, error) {
	groupMembers, resp, err := c.Groups.ListGroupMembers(key, &gitlab.ListGroupMembersOptions{})
	if err != nil {
		log.WithField("response", resp).WithError(err).Error("MailsFromGroupProject")
		if resp.StatusCode == 404 {
			metrics.WrongProjectCounter.Inc()
		}
		return nil, err
	}
	return groupMembers, nil
}

// func (c *Client) TestCallback(key string) (interface{}, error) {
// 	return []string{"test", "blabla", key}, nil
// 	// return nil, fmt.Errorf("error")
// }

// MailsFromGroupProject returns distincts mails from a project and its group
func (c *Client) MailsFromGroupProject(group, project string) ([]string, error) {
	const level = 40

	groupMembers, err := cache.GetWithCallback(group, c.GetGitlabGroupMembers)
	// groupMembers, err := cache.GetWithCallback(group, c.TestCallback)
	// groupMembers, err := groupMembersCache.GetWithCallback(group, c.GetGitlabGroupMembers)
	if err != nil && groupMembers == nil {
		return nil, err
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
