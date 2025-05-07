package jsonld

import (
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

// JSONLDRenderer is the renderer for JSON-LD fenced code blocks.
type JSONLDRenderer struct{}

// RegisterFuncs registers the renderer functions.
func (r *JSONLDRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindFencedCodeBlock, r.renderFencedCodeBlock)
}

// renderFencedCodeBlock renders the content of a JSON-LD block as a script tag.
func (r *JSONLDRenderer) renderFencedCodeBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	fmt.Printf("renderFencedCodeBlock: entering=%v, node=%T\n", entering, node)
	if !entering {
		return ast.WalkContinue, nil
	}

	// Check if the block is a JSON-LD fenced code block
	block, ok := node.(*ast.FencedCodeBlock)
	if !ok {
		return ast.WalkContinue, nil
	}

	language := string(block.Language(source))
	if language != "jsonld" {
		// let other renderers handle this block if it's not JSON-LD
		return ast.WalkContinue, nil
	}

	// Render the block as a JSON-LD script tag
	_, _ = w.WriteString(`<script type="application/ld+json">`)
	_, _ = w.Write(block.Text(source))
	_, _ = w.WriteString(`</script>`)

	return ast.WalkContinue, nil
}

// JSONLDExtension is the Goldmark extension for JSON-LD blocks.
type JSONLDExtension struct{}

// Extend adds the JSON-LD parser and renderer to Goldmark.
func (e *JSONLDExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(parser.WithBlockParsers(
		util.Prioritized(parser.NewFencedCodeBlockParser(), 100),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&JSONLDRenderer{}, 100),
	))
}

// NewJSONLDExtension creates a new JSON-LD extension.
func NewJSONLDExtension() goldmark.Extender {
	return &JSONLDExtension{}
}
