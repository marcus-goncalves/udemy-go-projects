package pork

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	CloneCmd = &cobra.Command{
		Use:   "clone",
		Short: "clone a github repository",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) <= 0 {
				log.Fatalln("you must provide a repository url")
			}

			if err := CloneRepository(args[0], ref, create); err != nil {
				log.Fatalln("error cloning repo:", err)
			}
		},
	}
	ref    string
	create bool
)

func CloneRepository(repository string, ref string, shouldCreate bool) error {
	repo, err := NewGHRepo(repository)
	if err != nil {
		return err
	}

	if err := repo.Clone(viper.GetString("location")); err != nil {
		return err
	}

	if err := repo.Checkout(ref, shouldCreate); err != nil {
		return err
	}

	fmt.Printf("Clone repo to: %s\n", repo.RepoDir)
	return nil
}

func init() {
	CloneCmd.PersistentFlags().StringVar(&ref, "ref", "main", "specific reference to check out")
	CloneCmd.PersistentFlags().BoolVar(&create, "create", false, "create the reference if it doesn't exists")
}
