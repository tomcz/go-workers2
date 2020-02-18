package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	hostAddress string
	port        string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goworkersctl",
	Short: "goworkersctl is the GoWorkers2 cli tool",
	Long: `
   ___|    \ \        /          |                  ___ \
  |      _ \\ \  \   / _ \   __| |  /  _ \  __| __|    ) |
  |   | (   |\ \  \ / (   | |      <   __/ |  \__ \   __/
 \____|\___/  \_/\_/ \___/ _|   _|\_\\___|_|  ____/ _____|

  goworkersctl is a cli tool that allows a user to easily get data from their
  GoWorkers2 instance.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&hostAddress, "a", "localhost", "Host address for a specific goworkers2 instance.")
	rootCmd.PersistentFlags().StringVar(&port, "p", "8080", "Port number for a specific goworkers2 instance.")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
