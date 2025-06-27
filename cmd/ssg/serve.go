package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

var port int

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the static site locally",
	Run: func(cmd *cobra.Command, args []string) {
		go watchSourceDirs()

		dir := "dist"
		addr := fmt.Sprintf(":%d", port)
		fmt.Printf("Serving %s at http://localhost%s/\n", dir, addr)
		fs := http.FileServer(http.Dir(dir))
		http.Handle("/", fs)
		if err := http.ListenAndServe(addr, nil); err != nil {
			fmt.Println("Server error:", err)
		}
	},
}

func watchSourceDirs() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Watcher error:", err)
		return
	}
	defer watcher.Close()

	dirs := []string{"content", "static", "templates"}
	for _, dir := range dirs {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() {
				watcher.Add(path)
			}
			return nil
		})
	}

	fmt.Println("Watching for changes in content/, static/, templates/")
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			fmt.Printf("[WATCHER] %s: %s\n", event.Op, event.Name)
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("[WATCHER ERROR]", err)
		}
	}
}

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to serve on")
}
