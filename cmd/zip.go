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

// zipCmd represents the zip command
var zipCmd = &cobra.Command{
	Use:   "zip",
	Short: "Method to zip a file / directory.",
	Long:  "Method to zip a file / directory.",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()
		err := services.Zip(args)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Done in ", time.Since(start).Seconds(), "Seconds")
	},
}

func init() {
	rootCmd.AddCommand(zipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
