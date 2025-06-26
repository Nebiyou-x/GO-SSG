package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the static site",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hey, I am Build")
	},
}
