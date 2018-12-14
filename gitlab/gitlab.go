package gitlab

import (
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/factorysh/minasan/metrics"
	log "github.com/sirupsen/logrus"
	gitlab "github.com/xanzy/go-gitlab"
)

// Client for Gitlab REST API
type Client struct {
	*gitlab.Client
}

// NewClientWithGitlabPrivateToken returns a new Client with a Gitlab's private token
func NewClientWithGitlabPrivateToken(client *http.Client, gitlabDomain string, privateToken string) *Client {
	gl := gitlab.NewClient(client, privateToken)
	gl.SetBaseURL("https://" + gitlabDomain + "/api/v4")
	return &Client{gl}
}

func NewClientFromEnv(client *http.Client) *Client {
	return NewClientWithGitlabPrivateToken(client, os.Getenv("GITLAB_DOMAIN"), os.Getenv("GITLAB_PRIVATE_TOKEN"))
}

func (c *Client) MailsFromGroupProject(group, project string) ([]string, error) {
	const level = 40
	// Works for gitlab 9, but documentation talks about https://docs.gitlab.com/ce/api/members.html#list-all-members-of-a-group-or-project-including-inherited-members
	// It doesn't work with curl + private token, and go-gitlab seems to not implement it
	groupMembers, resp, err := c.Groups.ListGroupMembers(group, &gitlab.ListGroupMembersOptions{})
	if err != nil {
		log.WithField("response", resp).WithError(err).Error("MailsFromGroupProject")
		if resp.StatusCode == 404 {
			metrics.WrongProjectCounter.Inc()
		}
		return nil, err
	}
	mails := make(map[string]interface{})
	for _, member := range groupMembers {
		if member.AccessLevel < level || member.State != "active" {
			continue
		}
		user, resp, err := c.Users.GetUser(member.ID)
		if err != nil {
			log.WithField("response", resp).WithError(err).Error("MailsFromGroupProject")
			return nil, err
		}
		if user.Email == "" {
			log.WithField("user", user).Warning("User with an empty email")
		}
		mails[user.Email] = true
	}

	id := strings.Join([]string{group, project}, "/")
	members, resp, err := c.ProjectMembers.ListProjectMembers(id, &gitlab.ListProjectMembersOptions{})
	if err != nil {
		log.WithField("response", resp).WithError(err).Error("MailsFromGroupProject")
		return nil, err
	}
	for _, member := range members {
		if member.AccessLevel < level || member.State != "active" {
			continue
		}
		user, resp, err := c.Users.GetUser(member.ID)
		if err != nil {
			log.WithField("response", resp).WithError(err).Error("MailsFromGroupProject")
			return nil, err
		}
		if user.Email == "" {
			log.WithField("user", user).Warning("User with an empty email")
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
