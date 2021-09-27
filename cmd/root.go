package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "pomo",
	Short: "Prints current running timer value.",
	Long:  "Pomo is a cli tool for starting/stopping a simple pomodoro timer.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		up := viper.GetString("pomo.up")
		if len(up) > 0 {
			endt, err := time.Parse(time.RFC3339, up)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to parse time: \n\t%v\n", err)
			}
			timeLeft := endt.Sub(time.Now()).Round(time.Second)
			emoji := viper.Get("pomo.emoji")
			if timeLeft < time.Second*30 && timeLeft%(time.Second*2) == 0 {
				fmt.Printf("%v %v\n", "⚠️ ", timeLeft)
			} else {
				fmt.Printf("%v %v\n", emoji, timeLeft)
			}
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/pomo/config.json)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Set default config based on os config dir.
		config, err := os.UserConfigDir()
		cobra.CheckErr(err)
		abspath, _ := os.Executable()
		x := path.Base(abspath)
		configDir := path.Join(config, x)
		cfgFile = path.Join(configDir, "config.json")
		viper.SetConfigFile(cfgFile)
	}

	viper.SetDefault("pomo.emoji", "🍅")
	viper.AutomaticEnv() // read in environment variables that match

	ensureConfig()
	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error with config: \n\t%v\n", err)
	}
}

func ensureConfig() {
	if cfgFile == "" {
		return
	}
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(cfgFile), 0700)
		os.WriteFile(cfgFile, []byte("{}"), 0600)
	}
}
