package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "this is a demo istio micro",
	Long:  `this is a demo istio micro ,has many`,
	Run: func(cmd *cobra.Command, args []string) {
		//  Do Stuff Here
	},
}

func cmdInit() {
	rootCmd.AddCommand(versionCmd)
}

func Execute() {
	cmdInit()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
