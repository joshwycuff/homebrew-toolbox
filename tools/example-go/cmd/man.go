package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var manCmd = &cobra.Command{
	Use:          "man",
	Short:        "Generate man pages",
	Hidden:       true,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		header := &doc.GenManHeader{
			Title:   "example-go",
			Section: "1",
		}
		return doc.GenManTree(rootCmd, header, ".")
	},
}

func init() {
	rootCmd.AddCommand(manCmd)
}
