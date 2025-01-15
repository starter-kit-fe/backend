package main

import (
	"admin/internal/app"
	"admin/internal/constant"
	"log"

	"github.com/spf13/cobra"
)

var (
	addr string
)
var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run app and more with a specified mode.",
	Run:   run,
}

func init() {
	cmdRun.Flags().StringVar(&addr, "addr", constant.PORT, "Running addr"+constant.PORT)
	rootCmd.AddCommand(cmdRun)
}

func run(cmd *cobra.Command, args []string) {
	// 初始化app
	application := app.NewApp(params)
	// application.Init()
	// 允许
	if err := application.Run(addr); err != nil {
		log.Fatalf("%v", err)
	}
}
