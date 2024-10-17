package cmd

import (
	"github.com/project-act/be-graph-local-go/gh"
	"github.com/spf13/cobra"
)

var (
	ghCmd = &cobra.Command{
		Use:     "gh [command]",
		Aliases: []string{"github", "g"},
		Short:   "GitHub API access.",
		Long:    `GitHub API access. Login using "be-graph-local-go gh login".`,
		Run: func(_ *cobra.Command, args []string) {
			gh.WillNorris()
			// fmt.Println("Hello from ghCmd!!!")
		},
	}
)
