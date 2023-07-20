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

type ForkResponse struct {
	CloneURL string `json:"clone_url"`
	FullName string `json:"full_name"`
}

var ForkCmd = &cobra.Command{
	Use:   "fork",
	Short: "fork a github repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("you must provide a repository")
		}

		if err := ForkRepository(args[0]); err != nil {
			log.Fatalln("unable to fork repository", err)
		}
	},
}

func ForkRepository(repository string) error {
	values := strings.Split(repository, "/")
	if len(values) != 2 {
		return fmt.Errorf("repository must be owner/project")
	}

	return GithubAPI().Call("fork", map[string]string{
		"owner": values[0],
		"repo":  values[1],
	})
}

func GetForkResouce() *nap.RestResource {
	forkRouter := nap.NewRouter()
	forkRouter.RegisterFunc(202, ForkSuccess)
	forkRouter.RegisterFunc(301, func(resp *http.Response, content interface{}) error {
		return fmt.Errorf("you must set an authentication token")
	})

	fork := nap.NewResource("/repos/{{.owner}}/{{.repo}}//forks", "POST", forkRouter)
	return fork
}

func ForkSuccess(resp *http.Response, _ interface{}) error {
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respContent := ForkResponse{}
	json.Unmarshal(content, &respContent)
	fmt.Printf("forked to: %s", respContent.FullName)

	return nil
}
