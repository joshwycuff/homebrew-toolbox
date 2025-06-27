package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/joshwycuff/homebrew-toolbox/tools/example-go/cmd/math"
)

var rootFlagInfo bool
var rootFlagDebug bool
var rootFlagTrace bool
var rootFlagVerbosity int

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
	writer := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	writer.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %s |", i))
	}
	log.Logger = log.Output(writer)

	if rootFlagTrace || rootFlagVerbosity >= 3 {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	} else if rootFlagDebug || rootFlagVerbosity >= 2 {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if rootFlagInfo || rootFlagVerbosity >= 1 {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	//log.Panic().Msg("panic")
	//log.Fatal().Msg("fatal")
	//log.Error().Msg("error")
	//log.Warn().Msg("warn")
	//log.Info().Msg("info")
	//log.Debug().Msg("debug")
	//log.Trace().Msg("trace")

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

	rootCmd.PersistentFlags().BoolVar(&rootFlagInfo, "info", false, "Enable debug level logging")
	rootCmd.PersistentFlags().BoolVar(&rootFlagDebug, "debug", false, "Enable debug level logging")
	rootCmd.PersistentFlags().BoolVar(&rootFlagTrace, "trace", false, "Enable trace level logging")
	rootCmd.PersistentFlags().CountVarP(&rootFlagVerbosity, "verbosity", "v", "Increase verbosity of logging")
}
