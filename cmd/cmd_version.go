package main

import (
	"os"
	"runtime"
	"admin/internal/constant"

	"github.com/spf13/cobra"
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print current version of admin",
	Run:   printVersion,
}
var nameOnly bool
var VERSION = "N/A"

func init() {
	cmdVersion.Flags().BoolVarP(&nameOnly, "name", "n", false, "print version name only")
	rootCmd.AddCommand(cmdVersion)
}

func printVersion(cmd *cobra.Command, args []string) {
	if nameOnly {
		os.Stdout.WriteString(constant.VERSION + "\n")
		return
	}
	version := "\n" + constant.NAME + " version: " + constant.VERSION + "\n"
	version += "Environment: " + runtime.Version() + " " + runtime.GOOS + "/" + runtime.GOARCH + "\n"
	os.Stdout.WriteString(version)

}
