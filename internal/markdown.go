package internal

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

func MarkdownToHTML(md string) string {
	renderer := html.NewRenderer(html.RendererOptions{})
	htmlBytes := markdown.ToHTML([]byte(md), nil, renderer)
	return string(htmlBytes)
}
