package math

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"strconv"
)

var addCmd = &cobra.Command{
	Use:          "add <number> [<number>...]",
	Short:        "Add N integers together.",
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: true,
	RunE:         add,
}

func add(cmd *cobra.Command, args []string) error {

	var total int64

	for _, arg := range args {
		log.Debug().Msg(fmt.Sprintf("Adding %s", arg))
		n, err := strconv.ParseInt(arg, 10, 64)
		if err != nil {
			return err
		}
		total += n
	}

	fmt.Printf("%d", total)

	return nil
}

func init() {
	Cmd.AddCommand(addCmd)
}
