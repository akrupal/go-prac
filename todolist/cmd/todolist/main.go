package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	description = "todolist server"
)

var (
	debug bool
)

var rootCmd = &cobra.Command{
	Use:              "todolist",
	Short:            description,
	Long:             ``,
	PersistentPreRun: doPersistantPreRun,
	SilenceErrors:    true,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug logging")
}

func doPersistantPreRun(cmd *cobra.Command, args []string) {
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Error().Err(err).Msg(description + "failed")
		os.Exit(1)
	}
}
