package main

import (
	"admin/internal/app"

	"github.com/spf13/cobra"
)

var cmdInit = &cobra.Command{
	Use:   "init",
	Short: "Init database",
	Run:   runinit,
}

func init() {
	rootCmd.AddCommand(cmdInit)
}

func runinit(cmd *cobra.Command, args []string) {
	application := app.NewApp(params)
	application.Init()

}
