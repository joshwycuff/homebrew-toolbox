package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"log/slog"
	"os"
)

var docCmd = &cobra.Command{
	Use:          "doc",
	Short:        "Generate markdown reference documentation for the project",
	Hidden:       true,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := os.MkdirAll("./docs", 0755); err != nil {
			slog.Error(fmt.Sprintf("Error creating docs directory: %v\n", err))
			return err
		}
		slog.Info("Generating markdown reference documentation for the project")
		return doc.GenMarkdownTree(rootCmd, "./docs")
	},
}

func init() {
	rootCmd.AddCommand(docCmd)
}
