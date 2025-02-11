package main

import (
	sqlitedb "todolist/pkg/db"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	bindAddress string
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Runs the Todolist Server",
	Long:  ``,
	RunE:  doServe,
}

// func init(){
// 	rootCmd
// }

func doServe(cmd *cobra.Command, args []string) error {
	log.Info().Msg(description + " starting")

	tododb, err := sqlitedb.CreateDb()
	if err != nil {
		return err
	}

	todostore := store.NewSQLStore(tododb)

}
