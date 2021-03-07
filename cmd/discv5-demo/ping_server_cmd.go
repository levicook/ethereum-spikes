package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/pkg/errors"
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
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return errors.WithStack(err)
	}

	db, err := enode.OpenDB("") // empty path returns an in-memory db
	if err != nil {
		return errors.WithStack(err)
	}

	udpAddr := net.UDPAddr{IP: net.IP{127, 0, 0, 1}}
	socket, err := net.ListenUDP("udp4", &udpAddr)
	if err != nil {
		return errors.WithStack(err)
	}

	cfg := discover.Config{
		PrivateKey: privateKey,
	}

	localNode := enode.NewLocalNode(db, privateKey)
	fmt.Println("localNode.ID()", localNode.ID())
	fmt.Println("localNode.Node().IP()", localNode.Node().IP())
	fmt.Println("localNode.Node().UDP()", localNode.Node().UDP())
	fmt.Println("localNode.Node().TCP()", localNode.Node().TCP())
	fmt.Println("localNode.Node().String()", localNode.Node().String())

	listenV5, err := discover.ListenV5(socket, localNode, cfg)
	if err != nil {
		return err
	}
	defer listenV5.Close()

	// wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Waiting for shutdown signals...")
	<-ch
	fmt.Println("Received signal, shutting down...")

	return nil
}
