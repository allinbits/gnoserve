package gnomark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gnolang/gno/gno.land/pkg/gnoweb"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

var (
	KindGnoMark     = ast.NewNodeKind("GnoMarkBlock")
	gnoMarkStartTag = []byte("<gno-mark>")
	gnoMarkEndTag   = []byte("</gno-mark>")

	templateRegistry = map[string]func(string) string{
		"frame":   gnoFrameRender,
		"html":    noHtmlMsg,
		"json+ld": renderJsonLd,
	}
)

type gnoMarkBlock struct {
	ast.BaseBlock
	Content string
}

var _ ast.Node = (*gnoMarkBlock)(nil)

func (b *gnoMarkBlock) Kind() ast.NodeKind {
	return KindGnoMark
}

func (b *gnoMarkBlock) Dump(source []byte, level int) {
	m := map[string]string{
		"Content": b.Content,
	}
	ast.DumpHelper(b, source, level, m, nil)
}

type gnoMarkParser struct{}

func (p *gnoMarkParser) Open(parent ast.Node, reader text.Reader, _ parser.Context) (ast.Node, parser.State) {
	_ = parent // REVIEW: any benefit to using parent?
	line, _ := reader.PeekLine()
	if !bytes.HasPrefix(line, gnoMarkStartTag) {
		return nil, parser.NoChildren
	}
	reader.AdvanceLine()
	return &gnoMarkBlock{}, parser.NoChildren
}

func (p *gnoMarkParser) Continue(node ast.Node, reader text.Reader, _ parser.Context) parser.State {
	line, _ := reader.PeekLine()
	if line == nil || bytes.HasPrefix(line, gnoMarkEndTag) {
		return parser.Close
	}
	block := node.(*gnoMarkBlock)
	block.Content += string(line)
	return parser.Continue
}

func (p *gnoMarkParser) Close(_ ast.Node, reader text.Reader, _ parser.Context) {
	for {
		line, _ := reader.PeekLine()
		if line == nil || bytes.HasPrefix(line, gnoMarkEndTag) {
			reader.AdvanceLine()
			break
		}
		reader.AdvanceLine()
	}
}

func (p *gnoMarkParser) CanInterruptParagraph() bool {
	return true
}

func (p *gnoMarkParser) CanAcceptIndentedLine() bool {
	return false
}

func (p *gnoMarkParser) Trigger() []byte {
	return []byte{'<'}
}

// gnoMarkRenderer renders the gnoMark block as HTML.
type gnoMarkRenderer struct {
	client *gnoweb.HTMLWebClient
}

type GnoMarkData struct {
	GnoMark string          `json:"gnoMark"`
	RawData json.RawMessage `json:"-"`
}

func (g *GnoMarkData) UnmarshalJSON(data []byte) error {
	var temp map[string]json.RawMessage
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if rawGnoMark, ok := temp["gnoMark"]; ok {
		if err := json.Unmarshal(rawGnoMark, &g.GnoMark); err != nil {
			return err
		}
	}

	g.RawData = data
	return nil
}

func (r *gnoMarkRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindGnoMark, r.renderGnoMarkBlock)
}

// XXX allow everything
func isValidFragment(content string) bool {
	return true
}

func (r *gnoMarkRenderer) renderGnoMarkBlock(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	_ = source // source with tags
	if !entering {
		return ast.WalkContinue, nil
	}

	b, ok := node.(*gnoMarkBlock)
	if !ok {
		return ast.WalkContinue, nil
	}

	content := strings.TrimSuffix(b.Content, "<gno-mark>") // FIXME: is this necessary?

	var gnoMarkData GnoMarkData
	err := gnoMarkData.UnmarshalJSON([]byte(content))
	if err != nil {
		if isValidFragment(content) {
			gnoMarkData.GnoMark = "html" // try to render as HTML (if enabled)
		} else {
			return ast.WalkStop, nil
		}
	} else {
		fmt.Printf("gnoMarkData: %s", gnoMarkData.RawData)
	}

	template, ok := templateRegistry[gnoMarkData.GnoMark]

	if !ok {
		return ast.WalkStop, nil
	}

	_, _ = w.WriteString(template(content))

	return ast.WalkContinue, nil
}
