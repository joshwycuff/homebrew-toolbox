package complete

import (
	"github.com/spf13/cobra"
)

func Directory(cmd *cobra.Command, args []string, toComplete string) ([]cobra.Completion, cobra.ShellCompDirective) {
	return nil, cobra.ShellCompDirectiveDefault
}
