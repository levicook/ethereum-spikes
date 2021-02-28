package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
	"github.com/spf13/cobra"
)

var pingServerCmd = &cobra.Command{
	Use:  "ping-server",
	RunE: runPingServer,
}

func init() {
	rootCmd.AddCommand(pingServerCmd)
}

func runPingServer(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	opts := libp2p.ChainOptions(
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"), // listens on a random local TCP port,
		libp2p.Ping(false), // without running the built-in ping protocol
	)

	host, err := libp2p.New(ctx, opts)
	if err != nil {
		return err
	}
	defer host.Close()

	// configure our own ping protocol
	pingService := &ping.PingService{Host: host}
	host.SetStreamHandler(ping.ID, pingService.PingHandler)

	// print the node's PeerInfo in multiaddr format
	hostInfo := peer.AddrInfo{
		ID:    host.ID(),
		Addrs: host.Addrs(),
	}
	hostAddrs, err := peer.AddrInfoToP2pAddrs(&hostInfo)
	if err != nil {
		return err
	}
	for _, addr := range hostAddrs {
		fmt.Println("host addr:", addr)
	}

	// wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Waiting for shutdown signals...")
	<-ch
	fmt.Println("Received signal, shutting down...")

	return nil
}
