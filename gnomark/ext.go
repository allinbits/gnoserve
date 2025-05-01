package gnomark

import (
	"github.com/gnolang/gno/gno.land/pkg/gnoweb"
	img64 "github.com/tenkoh/goldmark-img64"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
	"go.abhg.dev/goldmark/mermaid"
	"net/http"
)

// GnoMarkExtension is the Goldmark extension adding block parsers and renderers
// for GnoMark blocks: <gno-mark>...</gno-mark>
type GnoMarkExtension struct {
	Client *gnoweb.HTMLWebClient
}

func (e *GnoMarkExtension) Extend(m goldmark.Markdown) {
	(&mermaid.Extender{
		RenderMode:   0,
		Compiler:     nil,
		CLI:          nil,
		MermaidURL:   "",
		ContainerTag: "",
		NoScript:     false,
		Theme:        "",
	}).Extend(m) // mermaid

	// image embedding
	img64.Img64.Extend(m)
	m.Renderer().AddOptions(img64.WithFileReader(img64.AllowRemoteFileReader(http.DefaultClient)))

	// Setup Gnomark
	m.Parser().AddOptions(parser.WithBlockParsers(
		util.Prioritized(&gnoMarkParser{}, 500),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&gnoMarkRenderer{client: e.Client}, 500),
	))

	// Enable auto heading IDs for better linking
	m.Parser().AddOptions(parser.WithAutoHeadingID())
}
