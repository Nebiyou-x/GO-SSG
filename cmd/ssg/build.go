package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"go_proj/internal"

	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the static site",
	Run: func(cmd *cobra.Command, args []string) {
		srcDir := "content"
		outDir := "dist"
		err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if !strings.HasSuffix(info.Name(), ".md") {
				return nil
			}
			relPath, err := filepath.Rel(srcDir, path)
			if err != nil {
				return err
			}
			outPath := filepath.Join(outDir, strings.TrimSuffix(relPath, ".md")+".html")
			outDirPath := filepath.Dir(outPath)
			if err := os.MkdirAll(outDirPath, 0755); err != nil {
				return err
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			html := internal.MarkdownToHTML(string(data))
			title := strings.TrimSuffix(info.Name(), ".md")
			final, err := internal.ApplyTemplate(title, html)
			if err != nil {
				return err
			}
			if err := ioutil.WriteFile(outPath, []byte(final), 0644); err != nil {
				return err
			}
			fmt.Printf("Processed %s -> %s\n", path, outPath)
			return nil
		})
		if err != nil {
			fmt.Println("Build error:", err)
		} else {
			fmt.Println("Build complete.")
		}

		// Copy static assets
		staticDir := "static"
		err = filepath.Walk(staticDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				if os.IsNotExist(err) {
					return nil // skip if static dir doesn't exist
				}
				return err
			}
			relPath, err := filepath.Rel(staticDir, path)
			if err != nil {
				return err
			}
			outPath := filepath.Join(outDir, relPath)
			if info.IsDir() {
				return os.MkdirAll(outPath, 0755)
			}
			// Copy file
			in, err := os.Open(path)
			if err != nil {
				return err
			}
			defer in.Close()
			out, err := os.Create(outPath)
			if err != nil {
				return err
			}
			defer out.Close()
			_, err = io.Copy(out, in)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			fmt.Println("Static asset copy error:", err)
		} else {
			fmt.Println("Static assets copied.")
		}
	},
}
