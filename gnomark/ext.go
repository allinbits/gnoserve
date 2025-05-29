package gnomark

import (
	"github.com/gnolang/gno/gno.land/pkg/gnoweb"
	img64 "github.com/tenkoh/goldmark-img64"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/mermaid"
)

// templateRegistry maps GnoMark types to their rendering functions.
var templateRegistry = map[string]func(string) string{}

// NewGnoMarkExtension adds a new GnoMark code block extension to Goldmark.
func RegisterTemplate(name string, fn func(string) string) {
	if _, exists := templateRegistry[name]; exists {
		panic("template already registered: " + name)
	}
	templateRegistry[name] = fn
}

// GnoMarkExtension is the Goldmark extension adding block parsers and renderers
// for GnoMark blocks: ```gnomark {json} ```
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
	_ = img64.Img64
	//img64.Img64.Extend(m)
	//m.Renderer().AddOptions(img64.WithFileReader(img64.AllowRemoteFileReader(http.DefaultClient)))

	// Setup Gnomark
	NewGnoMarkExtension().Extend(m)

	// Enable auto heading IDs for better linking
	m.Parser().AddOptions(parser.WithAutoHeadingID())
}
