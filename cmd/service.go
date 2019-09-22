package cmd

import (
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Will run the service.",
	Run: func(cmd *cobra.Command, args []string) {
		serviceHelper("service")
	},
}

var serviceStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Issues the control function to start the service.",
	Run: func(cmd *cobra.Command, args []string) {
		serviceHelper("start")
	},
}

var serviceStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Issues the control function to stop the service.",
	Run: func(cmd *cobra.Command, args []string) {
		serviceHelper("stop")
	},
}

var serviceRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Issues the control function to restart the service.",
	Run: func(cmd *cobra.Command, args []string) {
		serviceHelper("restart")
	},
}

var serviceInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Issues the control function to install the service.",
	Run: func(cmd *cobra.Command, args []string) {
		serviceHelper("install")
	},
}

var serviceUninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Args:  cobra.ArbitraryArgs,
	Short: "Issues the control function to uninstall the service.",
	Run: func(cmd *cobra.Command, args []string) {
		serviceHelper("uninstall")
	},
}

func init() {
	// rootCmd.AddCommand(serviceCmd)
	// serviceCmd.PersistentFlags().BoolP("toggle", "t", false, "Help message for toggle")

	serviceCmd.AddCommand(serviceStartCmd)
	serviceCmd.AddCommand(serviceStopCmd)
	serviceCmd.AddCommand(serviceRestartCmd)
	serviceCmd.AddCommand(serviceInstallCmd)
	serviceCmd.AddCommand(serviceUninstallCmd)

}
