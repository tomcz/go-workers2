package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Go-Workers2 Stats info",
	Long: `Use the stats command to retrieve stats info from a specified host address and
	port number, like so:

	goworkersctl stats --a 127.0.0.1 --p 8080`,
	RunE: runGetStats,
}

func init() {
	rootCmd.AddCommand(statsCmd)
}

func runGetStats(cmd *cobra.Command, args []string) error {
	address := "http://" + hostAddress + ":" + port + "/stats"

	resp, err := http.Get(address)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Body: %v\n", string(body))
	return nil
}
