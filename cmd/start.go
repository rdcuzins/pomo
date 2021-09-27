package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a pomo timer. Defaults to 25m",
	Long:  "This will start a timer based on string provided. Example (30m20s) ",
	Run:   run,
}

func init() {
	rootCmd.AddCommand(startCmd)
	viper.SetDefault("pomo.duration", "25m")
}

func run(cmd *cobra.Command, args []string) {
	duration := "25m"
	if len(args) > 0 {
		duration = args[0]
	}
	dur, err := time.ParseDuration(duration)
	if err != nil {
		fmt.Fprintf(os.Stderr, "TimeParse error: \n\t%v\n", err)
	}
	up := time.Now().Add(dur).Format(time.RFC3339)
	viper.Set("pomo.up", up)
	if err := viper.WriteConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to write config: \n\t%v\n", err)
	}
}
