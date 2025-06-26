package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the static site locally",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hey, I am serve")
	},
}
