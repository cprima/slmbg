package cmd

import (
	"github.com/spf13/cobra"
)

var tickCmd = &cobra.Command{
	Use:   "tick",
	Short: "generates a fresh png file at current system time",
	Long: `With ./slmbg tick a png file is created depicting
the world map in daylight and nighttime
at the current system time (which is hopefully "now" ;):

For the math behind this and the Android app visit http://slm.prdv.de/ .`,
	Run: func(cmd *cobra.Command, args []string) {
		tick()
	},
}

func init() {
	serviceCmd.AddCommand(tickCmd)
}
