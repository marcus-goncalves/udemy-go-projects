package pork

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var DocsCmd = &cobra.Command{
	Use:   "docs",
	Short: "read documentation for a repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			log.Fatalln("you must provide a repository url")
		}
		content := GetRepositoryDoc(args[0])
		fmt.Println(content)
	},
}

func GetRepositoryDoc(repository string) string {
	return repository
}
