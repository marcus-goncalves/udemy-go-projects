package pork

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	nap "udemy-projects.com/devops-tools/nap/cmd"
)

type PullRequestPayload struct {
	Title        string `json:"title"`
	Message      string `json:"body"`
	SourceBranch string `json:"head"`
	DestBranch   string `json:"base"`
	Modify       bool   `json:"maintainer_can_modify"`
}

type PullRequestResponse struct {
	Url string `json:"html_url"`
}

var (
	destRepo         string
	sourceRepo       string
	pullRequestTitle string
	pullRequestMsg   string

	PullRequestCmd = &cobra.Command{
		Use:   "pullrequest",
		Short: "Create a Pull Request",
		Run: func(cmd *cobra.Command, args []string) {
			if err := CreatePullRequest(); err != nil {
				log.Fatalln("failed to create pull request:", err)
			}
		},
	}
)

func CreatePullRequest() error {
	sourceValues := strings.Split(sourceRepo, ":")
	if !(len(sourceValues) == 1 || len(sourceValues) == 2) {
		return fmt.Errorf("source repository must be [owner:]branch, but got %v", sourceRepo)
	}

	destBranchValues := strings.Split(destRepo, ":")
	if len(destBranchValues) != 2 {
		fmt.Printf("destination must be owner/project:branch, but got %v", destRepo)
	}

	destValues := strings.Split(destBranchValues[0], "/")
	if len(destValues) != 2 {
		fmt.Printf("destination must be owner/project:branch, but got %v", destRepo)
	}

	payload := &PullRequestPayload{
		Title:        pullRequestTitle,
		Message:      pullRequestMsg,
		SourceBranch: sourceRepo,
		DestBranch:   destRepo,
		Modify:       true,
	}
	return GithubAPI().Call("pullrequest", map[string]string{
		"owner":   destValues[0],
		"project": destValues[1],
	}, payload)
}

func PullRequestSuccess(resp *http.Response, _ interface{}) error {
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respContent := PullRequestResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Printf("created pull request:%s", respContent.Url)

	return nil
}

func PullRequestDefaultRouter(resp *http.Response, _ interface{}) error {
	return fmt.Errorf("status code %d", resp.StatusCode)
}

func GetPullRequestResource() *nap.RestResource {
	router := nap.NewRouter()
	router.RegisterFunc(201, PullRequestSuccess)
	router.DefaultRouter = PullRequestDefaultRouter
	resource := nap.NewResource("/repos/{{.owner}}/{{.project}}/pulls", "POST", router)

	return resource
}

func init() {
	PullRequestCmd.Flags().StringVarP(&sourceRepo, "source", "s", "", "source repository")
	PullRequestCmd.Flags().StringVarP(&destRepo, "destination", "d", "", "destination repository")
	PullRequestCmd.Flags().StringVarP(&pullRequestTitle, "title", "t", "", "basic pull request")
	PullRequestCmd.Flags().StringVarP(&pullRequestMsg, "message", "m", "", "pull request message")
}
