package cmd

import (
	"fmt"

	"github.com/cprior/slmbg/slmbglib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints version information.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
		fmt.Println("slmbglib Buildtime: ", slmbglib.Buildtime)
		fmt.Println("slmbglib Version: ", slmbglib.Version)
		fmt.Println("slmbglib Git hash: ", slmbglib.Githash)
		fmt.Println("config width: ", viper.GetInt("width"))
		fmt.Println("config DaylightImageFilename: ", viper.GetString("DaylightImageFilename"))
		fmt.Println("config NighttimeImageFilename: ", viper.GetString("NighttimeImageFilename"))
		fmt.Println("config OutputImageFilename: ", viper.GetString("OutputImageFilename"))
	},
}

func init() {
	serviceCmd.AddCommand(versionCmd)
}
