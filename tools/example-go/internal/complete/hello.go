package complete

import (
	"github.com/spf13/cobra"
	"os"
)

func Hello(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	user := os.Getenv("USER")

	if user == "" {
		return []string{"world"}, cobra.ShellCompDirectiveNoFileComp
	}
	return []string{"world", user}, cobra.ShellCompDirectiveNoFileComp
}
