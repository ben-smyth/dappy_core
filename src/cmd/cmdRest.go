/*
Copyright © 2023 NAME HERE ben.df.smyth@gmail.com

*/
package cmd

import (
	"dappy_core/rest"
	"fmt"

	"github.com/spf13/cobra"
)

var restCmd = &cobra.Command{
	Use:   "rest",
	Short: "Run the RestAPI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rest called")
		rest.Run()
	},
}

func restApi() {

}

func init() {
	rootCmd.AddCommand(restCmd)

}
