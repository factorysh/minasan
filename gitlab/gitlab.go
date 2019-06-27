package gitlab

import (
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/factorysh/minasan/cache"
	"github.com/factorysh/minasan/metrics"
	log "github.com/sirupsen/logrus"
	gitlab "github.com/xanzy/go-gitlab"
)

// Client for Gitlab REST API
type Client struct {
	*gitlab.Client
	cache *cache.Cachedb
}

var gclient *Client = nil

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

// NewClientWithGitlabPrivateToken returns a new Client with a Gitlab's private token
func NewClientWithGitlabPrivateToken(client *http.Client, gitlabDomain string,
	privateToken string, ttl time.Duration, cachePath string) (*Client, error) {
	if gclient == nil {
		c, err := cache.New(ttl, cachePath)
		if err != nil {
			return nil, err
		}
		gl := gitlab.NewClient(client, privateToken)
		gl.SetBaseURL("https://" + gitlabDomain + "/api/v4")
		gclient = &Client{gl, c}
	}
	return gclient, nil
}

// NewClientFromEnv returns a new Client from environments
func NewClientFromEnv(client *http.Client) (*Client, error) {
	return NewClientWithGitlabPrivateToken(client, os.Getenv("GITLAB_DOMAIN"),
		os.Getenv("GITLAB_PRIVATE_TOKEN"), 5*time.Minute, "/tmp/minasan.db")
}

// MailsFromGroupProject returns distincts mails from a project and its group
func (c *Client) MailsFromGroupProject(group, project, lastChanceMail string) ([]string, error) {
	const level = 40

	groupMembers, err := c.cache.LazyGet(group, c.GetGitlabGroupMembers)
	if err != nil && groupMembers == nil {
		// Gitlab is unavailable, send a last chance email
		if lastChanceMail != "" {
			email := []string{lastChanceMail}
			log.Info("Sending last chance email")
			return email, nil
		}
		log.Warning("last_chance_mail is not set")
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
