/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"api-server/apiHandler"
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
		port, _ := cmd.Flags().GetInt("port")
		apiHandler.RunServer(port)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().Int("port", 8081, "Port to listen on")
}
