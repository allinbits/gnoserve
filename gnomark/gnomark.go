package gnomark

import (
	"encoding/json"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

var templateRegistry = map[string]func(string) string{
	"frame": gnoFrameRender,
	"html":  noHtmlMsg,
}

// gnoMarkRenderer renders the gomark fenced code block as HTML.
type gnoMarkRenderer struct{}

func (r *gnoMarkRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, r.renderGnoMarkBlock)
}

// getGnoMarkType extracts the gnoMark type from the JSON content.
func getGnoMarkType(jsonContent string) string {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonContent), &data); err == nil {
		if gnoMarkType, ok := data["gnoMark"].(string); ok {
			return gnoMarkType
		}
	}
	return "html"
}

func (r *gnoMarkRenderer) renderGnoMarkBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}

	block, ok := node.(*ast.FencedCodeBlock)
	if !ok {
		return ast.WalkContinue, nil
	}

	language := string(block.Language(source))
	if language != "gnomark" {
		return ast.WalkContinue, nil
	}
	content := string(block.Text(source))
	gnoMarkType := getGnoMarkType(content)
	template, ok := templateRegistry[gnoMarkType]
	if !ok {
		return ast.WalkStop, nil
	}

	_, _ = w.WriteString(template(content))
	return ast.WalkContinue, nil
}

// gnoMarkExtension is the Goldmark extension for gomark fenced code blocks.
type gnoMarkExtension struct{}

func (e *gnoMarkExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithBlockParsers(
		util.Prioritized(parser.NewFencedCodeBlockParser(), 100),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&gnoMarkRenderer{}, 100),
	))
}

// NewGnoMarkExtension creates a new gomark extension.
func NewGnoMarkExtension() *gnoMarkExtension {
	return &gnoMarkExtension{}
}
