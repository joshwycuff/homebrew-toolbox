package cmd

import (
	"fmt"
	"github.com/joshwycuff/homebrew-toolbox/tools/ff/internal/complete"
	"github.com/joshwycuff/homebrew-toolbox/tools/ff/internal/model"
	"golang.org/x/term"
	"os"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootFlagInitialQuery string
var rootFlagDirectory string
var rootFlagInfo bool
var rootFlagDebug bool
var rootFlagTrace bool
var rootFlagVerbosity int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "ff",
	Short:             "A TUI to search files by name or content",
	SilenceUsage:      true,
	CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	PersistentPreRunE: setupDefaultLogger,
	RunE:              run,
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

func run(cmd *cobra.Command, args []string) error {

	err := os.Chdir(rootFlagDirectory)
	if err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", rootFlagDirectory, err)
	}

	options := []tea.ProgramOption{
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	}

	if !term.IsTerminal(int(os.Stdout.Fd())) {
		ttyIn, ttyOut, err := forceTTY()
		if err != nil {
			return fmt.Errorf("failed to open /dev/tty: %w", err)
		}
		defer ttyIn.Close()
		defer ttyOut.Close()
		options = append(options, tea.WithInput(ttyIn), tea.WithOutput(ttyOut))
	}

	p := tea.NewProgram(
		model.New(rootFlagInitialQuery),
		options...,
	)

	finalModel, err := p.Run()
	if err != nil {
		return err
	}

	myModel, ok := finalModel.(model.Model)
	if !ok {
		return fmt.Errorf("final model is not of type model.Model")
	}

	if myModel.ExitCode == 0 {
		for _, selected := range myModel.Selected {
			fullPath := filepath.Join(rootFlagDirectory, selected)
			fmt.Println(fullPath)
		}
	} else {
		os.Exit(myModel.ExitCode)
	}

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
	rootCmd.PersistentFlags().StringVarP(&rootFlagInitialQuery, "initial-query", "i", "rg -l ", "Initial search query")
	rootCmd.PersistentFlags().StringVarP(&rootFlagDirectory, "directory", "d", ".", "Directory to search for files")
	rootCmd.PersistentFlags().BoolVar(&rootFlagInfo, "info", false, "Enable debug level logging")
	rootCmd.PersistentFlags().BoolVar(&rootFlagDebug, "debug", false, "Enable debug level logging")
	rootCmd.PersistentFlags().BoolVar(&rootFlagTrace, "trace", false, "Enable trace level logging")
	rootCmd.PersistentFlags().CountVarP(&rootFlagVerbosity, "verbosity", "v", "Increase verbosity of logging")

	err := rootCmd.RegisterFlagCompletionFunc("directory", complete.Directory)
	if err != nil {
		log.Error().Err(err).Msg("Failed to register flag completion function for 'directory'")
		os.Exit(1)
	}
}

func forceTTY() (*os.File, *os.File, error) {
	ttyIn, err := os.OpenFile("/dev/tty", os.O_RDONLY, 0)
	if err != nil {
		return nil, nil, err
	}
	ttyOut, err := os.OpenFile("/dev/tty", os.O_WRONLY, 0)
	if err != nil {
		ttyIn.Close()
		return nil, nil, err
	}
	return ttyIn, ttyOut, nil
}
