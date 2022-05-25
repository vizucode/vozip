/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/visstars7/vozip/services"
)

// unzipCmd represents the unzip command
var unzipCmd = &cobra.Command{
	Use:   "unzip",
	Short: "Unzip command.",
	Long:  "Command to unzip a zip file.",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		err := services.Unzip(args)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Done in ", time.Since(start).Seconds(), "Seconds")
	},
	Version: "1.0.0",
}

func init() {
	rootCmd.AddCommand(unzipCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unzipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unzipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
