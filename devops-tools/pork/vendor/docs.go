package pork

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	nap "udemy-projects.com/devops-tools/nap/cmd"
)

type REadmeResponse struct {
	Content string `json:"content"`
}

var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "read documentation for a repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("you must provide a repository url")
		}
		if err := GetRepositoryDoc(args[0]); err != nil {
			log.Fatalln("failed to get docs: ", err)
		}
	},
}

func ReadmeSuccess(resp *http.Response, _ interface{}) error {
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respContent := REadmeResponse{}
	json.Unmarshal(content, &respContent)
	buff, err := base64.StdEncoding.DecodeString(respContent.Content)
	if err != nil {
		return err
	}

	fmt.Println(string(buff))
	return nil
}

func GetRepositoryDoc(repository string) error {
	values := strings.Split(repository, "/")
	return GithubAPI().Call("docs", map[string]string{
		"owner":   values[0],
		"project": values[1],
	}, nil)
}

func ReadmeDefaultRouter(resp *http.Response, _ interface{}) error {
	return fmt.Errorf("status code: %d", resp.StatusCode)
}

func GetReadmeResource() *nap.RestResource {
	router := nap.NewRouter()
	router.RegisterFunc(200, ReadmeSuccess)
	router.DefaultRouter = ReadmeDefaultRouter
	resource := nap.NewResource("/repos/{{.owner}}/{{.project}}/readme", "GET", router)

	return resource
}
