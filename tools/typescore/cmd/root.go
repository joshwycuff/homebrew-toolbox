package cmd

import (
	"bufio"
	"fmt"
	"github.com/joshwycuff/homebrew-toolbox/tools/typescore/internal/score"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootFlagOnlyScores bool
var rootFlagArgDelimiter string

var rootCmd = &cobra.Command{
	Use:               "typescore",
	Short:             "Simple tool to score the typing difficulty of text",
	CompletionOptions: cobra.CompletionOptions{HiddenDefaultCmd: true},
	Args:              cobra.ArbitraryArgs,
	RunE:              run,
}

func run(cmd *cobra.Command, args []string) error {
	var scanner *bufio.Scanner
	if len(args) > 0 {
		input := strings.Join(args, rootFlagArgDelimiter)
		scanner = bufio.NewScanner(strings.NewReader(input))
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	for scanner.Scan() {
		line := scanner.Text()
		s := score.Score(line)
		if rootFlagOnlyScores {
			fmt.Printf("%d", s)
		} else {
			fmt.Printf("%s\t%d", line, s)
		}
		if scanner.Err() == nil {
			fmt.Print("\n")
		}
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
	rootCmd.Flags().BoolVarP(&rootFlagOnlyScores, "only-scores", "o", false, "Only print scores, no input line")
	rootCmd.Flags().StringVar(&rootFlagArgDelimiter, "arg-delimiter", " ", "Delimiter to insert between args")
}
