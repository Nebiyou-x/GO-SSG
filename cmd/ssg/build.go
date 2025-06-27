package main

import (
	"fmt"
	"io/ioutil"

	"go_proj/internal"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the static site",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadFile("sample.md")
		if err != nil {
			fmt.Println("Error reading sample.md:", err)
			return
		}
		html := internal.MarkdownToHTML(string(data))
		final, err := internal.ApplyTemplate("Sample Title", html)
		if err != nil {
			fmt.Println("Error applying template:", err)
			return
		}
		fmt.Println("Generated HTML:\n", final)
	},
}
