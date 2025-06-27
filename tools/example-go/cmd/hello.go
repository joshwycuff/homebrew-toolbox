package cmd

import (
	"fmt"
	"github.com/joshwycuff/homebrew-toolbox/tools/example-go/internal/complete"
	"strings"

	"github.com/spf13/cobra"
)

var helloFlagLoud bool

var helloCmd = &cobra.Command{
	Use:               "hello",
	Short:             "Say hello.",
	Args:              cobra.MaximumNArgs(1),
	ValidArgsFunction: complete.Hello,
	SilenceUsage:      true,
	Run:               hello,
}

func hello(cmd *cobra.Command, args []string) {
	greeting := "hello, "

	if len(args) > 0 {
		greeting = greeting + args[0]
	} else {
		greeting = greeting + "world"
	}

	if helloFlagLoud {
		greeting = strings.ToUpper(greeting) + "!!!"
	}

	fmt.Println(greeting)
}

func init() {
	rootCmd.AddCommand(helloCmd)

	helloCmd.Flags().BoolVar(&helloFlagLoud, "loud", false, "Make the greeting louder.")
}
