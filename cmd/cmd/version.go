package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xiaomeng79/istio-micro/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "this is a version",
	Long:  `this is a version explain`,
	Run: func(cmd *cobra.Command, args []string) {
		version.Ver()
	},
}
