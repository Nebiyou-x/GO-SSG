# Go Static Site Generator (SSG)

A simple, modern static site generator written in Go. Easily convert Markdown files to beautiful HTML, serve locally, and watch for changes.

## Features
| Feature                        | Description                                                                 |
|--------------------------------|-----------------------------------------------------------------------------|
| Markdown to HTML               | Converts all `.md` files in `content/` (including subfolders) to HTML        |
| Static Asset Copy              | Copies everything from `static/` to `dist/` (CSS, images, JS, etc.)         |
| HTML Templating                | Wraps content in a modern, customizable HTML template                        |
| Live Local Server              | `ssg serve` starts a local HTTP server with file watching                    |
| Auto Rebuild                   | Edits in `content/`, `static/`, or `templates/` trigger a rebuild           |
| Modern CSS                     | Beautiful, colorful default styles (easy to customize)                       |
| Configurable Port              | Serve on any port with `--port` flag                                        |
| Beginner Friendly              | Simple CLI, clear structure, and easy to extend                              |

---

## Quick Start

### 1. Clone the repository
```sh
git clone <https://github.com/Nebiyou-x/GO-SSG>
cd <https://github.com/Nebiyou-x/GO-SSG>/Go_proj
```

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
- Add your own HTML templates in `templates/` (if supported).

---

## Example Output

After running `./ssg build`, your `dist/` folder might look like:
```
dist/
├── index.html
├── posts/
│   └── post1.html
├── css/
│   └── style.css
└── img/
    └── logo.png
```

---

## Tips & Customization
- **Add a logo:** Place an image in `static/img/` and reference it in your Markdown or template.
- **Change colors:** Edit `static/css/style.css` for instant style changes.
- **Add JS:** Place scripts in `static/js/` and reference them in your Markdown or template.
- **Advanced:** Fork and extend the Go code for new features (pagination, blog index, etc.).

---

## How it Works
- **Build:** Recursively scans `content/` for Markdown files, converts them to HTML, and writes to `dist/` preserving the directory structure. Copies all static assets from `static/` to `dist/`.
- **Serve:** Starts a local HTTP server serving the `dist/` directory. Watches for changes in `content/`, `static/`, and `templates/` and automatically rebuilds the site.

---

## Troubleshooting
- If you see errors about missing Go modules, run `go mod tidy`.
- If you change Go code, rebuild with `go build -o ssg ./cmd/ssg`.
- If you get a permission error on `./ssg`, run `chmod +x ssg`.
