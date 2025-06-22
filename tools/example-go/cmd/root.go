package cmd

import (
	"github.com/joshwycuff/homebrew-toolbox/tools/example-go/cmd/math"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

var rootFlagDebug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "example-go",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	PersistentPreRunE: setupDefaultLogger,
}

func setupDefaultLogger(cmd *cobra.Command, args []string) error {
	var level slog.Level
	if rootFlagDebug {
		level = slog.LevelDebug
	} else {
		level = slog.LevelInfo
	}

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})
	slog.SetDefault(slog.New(handler))

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(math.Cmd)

	rootCmd.PersistentFlags().BoolVarP(&rootFlagDebug, "debug", "d", false, "Enable debug mode")
}
