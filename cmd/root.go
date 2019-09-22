package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/cprior/slmbg/images/icon"
	"github.com/cprior/slmbg/sunlightmap"
	"github.com/getlantern/systray"
	"github.com/reujab/wallpaper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "slmbg",
	Short: "Shows a day- and night map of the planet",
	Long: `The standalone companion of the Android app available at
Shows a day- and night map of the planet
http://slm.prdv.de/ .

Combines two png picture files of the day and night,
calculates an approximisation of twilight,
and may be used as the basis for wallpapers.`,
	Run: func(cmd *cobra.Command, args []string) {
		// systray.Run(onReady, onExit)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// go systray.Run(onReady, onExit)
	if err := serviceCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//called in the rootCmd by systray.Run(onReady, onExit)
func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Sunlightmap")
	systray.SetTooltip("slmbg")
	// go updateBackground()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		// fmt.Println("Requesting quit")
		systray.Quit()
		// fmt.Println("Finished quitting")
	}()
}
func onExit() {
	// serviceHelper("stop")
	os.Exit(1)
}

func tick() {
	slm := sunlightmap.NewStatic(viper.GetInt("width"), time.Now().Local())
	slm.DaylightImageFilename = viper.GetString("DaylightImageFilename")
	slm.NighttimeImageFilename = viper.GetString("NighttimeImageFilename")
	if viper.GetString("center") == "asia" {
		slm.CenterLongitude = 60
	} else if viper.GetString("center") == "foobar" {
		slm.CenterLongitude = 20
	} else {
		slm.CenterLongitude = 0
	}
	_ = sunlightmap.WriteStaticPng(&slm, viper.GetString("OutputImageFilename"))
	viper.Set("last_run", time.Now().Local().Format("2006-01-02 15:04:05"))
	_ = viper.WriteConfig()
}

func updateBackground() {
	return
	// for {
	// 	tick()
	// 	time.Sleep(time.Duration(viper.GetInt64("interval") * int64(time.Second)))
	// }
}

func init() {
	cobra.OnInitialize(initConfig)
	wallpaper.SetFromFile(viper.GetString("OutputImageFilename"))
	serviceCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// // Find home directory.
		// home, err := homedir.Dir()
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		// Search config in home directory with name ".slmbg" (without extension).
		// viper.AddConfigPath("/etc/slmbg")
		// viper.AddConfigPath(home + "/.slmbg")
		viper.AddConfigPath(".")
		viper.AddConfigPath("C:\\Users\\CPM\\go\\src\\github.com\\cprior\\slmbg")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
