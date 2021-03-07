package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var echoServerCmd = &cobra.Command{
	Use:  "echo-server",
	RunE: runEchoServer,
}

func init() {
	rootCmd.AddCommand(echoServerCmd)
}

func runEchoServer(cmd *cobra.Command, args []string) error {
	sig := make(chan os.Signal, 1)

	ctx, cancel := context.WithCancel(cmd.Context())

	go func(ctx context.Context) {
		conn, err := net.ListenPacket("udp", ":10000")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		var buf [256]byte
		for {
			select {
			default:
				n, _, err := conn.ReadFrom(buf[:])
				if err != nil && err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				s := fmt.Sprintf("%s", buf[:n])

				fmt.Printf("%s", s)

				if s == "quit\n" {
					fmt.Println("bye!")
					sig <- syscall.SIGINT
				}

			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	// wait for a SIGINT or SIGTERM signal
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Waiting for shutdown signals...")
	<-sig
	fmt.Println("Received signal, shutting down...")
	cancel()

	return nil
}
