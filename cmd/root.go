package cmd

import (
	"fire-press/apiservice/controller"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:  "root",
	Long: "default command",
	Run: func(cmd *cobra.Command, args []string) {
		//snowflake.InitSnowFlake()
		controller.InitController()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
