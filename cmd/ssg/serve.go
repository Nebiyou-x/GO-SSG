package main

import (
	"fmt"
	"go_proj/internal"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

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

	var mu sync.Mutex
	var timer *time.Timer
	debounce := func() {
		mu.Lock()
		defer mu.Unlock()
		if timer != nil {
			timer.Stop()
		}
		timer = time.AfterFunc(500*time.Millisecond, func() {
			fmt.Println("[WATCHER] Change detected, rebuilding...")
			buildSite()
		})
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			fmt.Printf("[WATCHER] %s: %s\n", event.Op, event.Name)
			debounce()
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("[WATCHER ERROR]", err)
		}
	}
}

func buildSite() {
	srcDir := "content"
	outDir := "dist"
	// Remove dist directory before rebuilding
	os.RemoveAll(outDir)
	// Reuse the build logic from buildCmd
	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
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
	// Copy static assets
	staticDir := "static"
	filepath.Walk(staticDir, func(path string, info os.FileInfo, err error) error {
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
	fmt.Println("Build complete.")
}

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to serve on")
}
