package pork

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SearchCmd = &cobra.Command{
	Use:   "search",
	Short: "search for github repositories by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		repositoryList := SearchByKeyword(args)
		for _, k := range repositoryList {
			fmt.Println(k)
		}
	},
}

func SearchByKeyword(keywords []string) []string {
	return []string{"myrepository"}
}
