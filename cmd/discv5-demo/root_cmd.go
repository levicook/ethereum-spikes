package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "discv5-demo",
	RunE: func(*cobra.Command, []string) error {
		fmt.Println("discv5 demos will hang off of here")
		return nil
	},
}
