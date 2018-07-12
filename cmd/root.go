// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
	Short: "shows a day- and night map of the planet",
	Long: `The standalone companion of the Android app available at
Shows a day- and night map of the planet
http://slm.prdv.de/ .

Combines two png picture files of the day and night,
calculates an approximisatin of twilight,
and may be used as the basis for wallpapers.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		systray.Run(onReady, onExit)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Sunlightmap")
	systray.SetTooltip("slmbg")
	go updateBackground()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}
func onExit() {}

func updateBackground() {
	for {
		slm := sunlightmap.NewStatic(viper.GetInt("width"), time.Now().Local())
		slm.DaylightImageFilename = viper.GetString("DaylightImageFilename")
		slm.NighttimeImageFilename = viper.GetString("NighttimeImageFilename")
		_ = sunlightmap.WriteStaticPng(&slm, viper.GetString("OutputImageFilename"))
		wallpaper.SetFromFile(viper.GetString("OutputImageFilename"))
		viper.Set("last_run", time.Now().Local().Format("2006-01-02 15:04:05"))
		viper.Set("center", "asia")
		//viper.Set("foo.bar", "baz")
		//viper.Set("ene", []string{"mene", "mu"})
		_ = viper.WriteConfig()
		time.Sleep(1000 * 30 * time.Millisecond)
		//fmt.Println(".")
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
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
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
