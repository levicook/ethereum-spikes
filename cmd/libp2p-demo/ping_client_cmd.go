package main

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"github.com/multiformats/go-multiaddr"
	"github.com/spf13/cobra"
)

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
	ctx := cmd.Context()

	opts := libp2p.ChainOptions(
		libp2p.Ping(false), // without running the built-in ping protocol
	)

	host, err := libp2p.New(ctx, opts)
	if err != nil {
		return err
	}
	defer host.Close()

	peerMultiAddr, err := multiaddr.NewMultiaddr(pingClientCfg.peerAddr)
	if err != nil {
		return err
	}
	peerAddrInfo, err := peer.AddrInfoFromP2pAddr(peerMultiAddr)
	if err != nil {
		return err
	}
	if err := host.Connect(ctx, *peerAddrInfo); err != nil {
		return err
	}

	fmt.Println("pinging", peerMultiAddr)

	pingService := &ping.PingService{Host: host}
	results := pingService.Ping(ctx, peerAddrInfo.ID)

	for i := 0; i < pingClientCfg.count; i++ {
		result := <-results
		if result.Error != nil {
			fmt.Println("pinged ", peerMultiAddr, "rtt:", result.RTT, "err:", result.Error)
		} else {
			fmt.Println("pinged ", peerMultiAddr, "rtt:", result.RTT)
		}
	}

	return nil
}
