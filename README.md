# Go Static Site Generator (SSG)

A simple, modern static site generator written in Go. Easily convert Markdown files to beautiful HTML, serve locally, and watch for changes.

## Features
- Convert Markdown files in `content/` (including subdirectories) to HTML in `dist/`
- Copy static assets from `static/` to `dist/`
- Basic HTML templating
- Live local server with file watching and auto-rebuild
- Modern, colorful default CSS
- Automatic rebuild on file changes
- Serve on a configurable port (default: 8080)

## Quick Start

### 1. Clone the repository


### 2. Install Go (if you don't have it)
- Download and install from: https://go.dev/dl/

### 3. Install dependencies
```sh
go mod tidy
```

### 4. Build the CLI
```sh
go build -o ssg ./cmd/ssg
```

### 5. Project Structure
```
content/         # Your Markdown files (can have subfolders)
static/          # Static assets (css, images, js, ...)
templates/       # (Optional) HTML templates
cmd/ssg/         # CLI source code
internal/        # Internal Go packages
dist/            # Output directory (auto-generated)
```

### 6. Add your content
- Place Markdown files in `content/` (e.g., `content/index.md`, `content/posts/post1.md`)
- Place static files in `static/` (e.g., `static/css/style.css`)

### 7. Build the site
```sh
./ssg build
```
- Output will be in the `dist/` directory, mirroring your `content/` and `static/` structure.

### 8. Serve locally with live reload
```sh
./ssg serve
```
- Visit [http://localhost:8080](http://localhost:8080) in your browser.
- Edit files in `content/`, `static/`, or `templates/` and the site will auto-rebuild.
- To use a different port: `./ssg serve --port 9000`

### 9. Customizing
- Edit `static/css/style.css` for your own styles.
- Edit Go code in `cmd/ssg/` and `internal/` to add features.

## How it Works
- **Build:** Recursively scans `content/` for Markdown files, converts them to HTML, and writes to `dist/` preserving the directory structure. Copies all static assets from `static/` to `dist/`.
- **Serve:** Starts a local HTTP server serving the `dist/` directory. Watches for changes in `content/`, `static/`, and `templates/` and automatically rebuilds the site.

## Troubleshooting
- If you see errors about missing Go modules, run `go mod tidy`.
- If you change Go code, rebuild with `go build -o ssg ./cmd/ssg`.
- If you get a permission error on `./ssg`, run `chmod +x ssg`.

## License
