package cmd

import (
	"fmt"

	"github.com/cprior/slmbg/slmbglib"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version information.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
		fmt.Println("slmbglib Buildtime: ", slmbglib.Buildtime)
		fmt.Println("slmbglib Version: ", slmbglib.Version)
		fmt.Println("slmbglib Git hash: ", slmbglib.Githash)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
