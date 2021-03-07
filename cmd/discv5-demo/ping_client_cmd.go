package main

import "github.com/spf13/cobra"

var pingClientCmd = &cobra.Command{
	Use:  "ping-client",
	RunE: runPingClient,
}

var pingClientCfg struct {
	peerAddr string
	count    int
}

func init() {
	rootCmd.AddCommand(pingClientCmd)

	pingClientCmd.Flags().StringVarP(&pingClientCfg.peerAddr, "peer-addr", "p", "", "peer address to ping (required)")
	if err := pingClientCmd.MarkFlagRequired("peer-addr"); err != nil {
		panic(err) // TODO(levi) there has to be a better way to handle this err
	}

	pingClientCmd.Flags().IntVarP(&pingClientCfg.count, "count", "c", 10, "number of pings to send before stopping")
}

func runPingClient(cmd *cobra.Command, args []string) error {
	return nil
}
