package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var retriesCmd = &cobra.Command{
	Use:   "retries",
	Short: "Go-Workers2 Retries info",
	Long: `Use the retries command to retrieve retries info from a specified host address and
	port number, like so:

	goworkersctl stats --a 127.0.0.1 --p 8080`,
	RunE: runGetRetries,
}

func init() {
	rootCmd.AddCommand(retriesCmd)
}

func runGetRetries(cmd *cobra.Command, args []string) error {
	address := "http://" + hostAddress + ":" + port + "/retries"

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
