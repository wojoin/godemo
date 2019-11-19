package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	verbose bool
	author  string
	license string
	rootCmd = &cobra.Command{
		Use: "cobra",
		Short: "Cobra is a library providing a simple interface that create" +
			"powerful modern CLI interfaces similar to git & to gools.",
		Long: "cobra is a library providing a simple interface that create " +
			"powerful modern CLI interfaces similar to git & to gools. It includes three concepts:\r\n" +
			"\t1) Commands\r\n\t2) Args\r\n\t3) Flags\r\n",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("verbose : ", verbose)
			fmt.Println("author: ", author)
			fmt.Println("cobra CLI app")
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&author, "author", "a", "join", "author name for copyright ")
	rootCmd.PersistentFlags().StringVarP(&license, "license", "l", "", "name of license for th eproject")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose info")
	rootCmd.Flags().Bool("version", true, "CLI version")
	rootCmd.Version = HC_VERSION
	//rootCmd.PersistentFlags().Bool("verbose", false, "verbose info")
}

// Execute executes root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
