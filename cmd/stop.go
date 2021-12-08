package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops active pomo timer if there is one.",
	Long:  "Stops timer by setting empty value in config parameter.",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("up", "")
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
