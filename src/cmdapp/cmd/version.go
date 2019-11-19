package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const HC_VERSION string = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print current version number of CLI hc",
	Long:  "All software has version. This is CLI hc version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(HC_VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
