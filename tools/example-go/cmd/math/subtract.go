package math

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var subtractCmd = &cobra.Command{
	Use:          "subtract <left> <right>",
	Aliases:      []string{"sub"},
	Short:        "Subtract one integer from another.",
	Args:         cobra.ExactArgs(2),
	SilenceUsage: true,
	RunE:         subtract,
}

func subtract(cmd *cobra.Command, args []string) error {

	if len(args) != 2 {
		return fmt.Errorf("expected 2 arguments, got %d", len(args))
	}

	left, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return err
	}

	right, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return err
	}

	result := left - right

	fmt.Printf("%d", result)

	return nil
}

func init() {
	Cmd.AddCommand(subtractCmd)
}
