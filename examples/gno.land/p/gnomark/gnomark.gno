package gnomark

import (
	"std"
	"strings"
	"time"
)

type Frame = func(path string, props ...interface{}) string

type WebHost struct {
	Base string
	Tag  string
	Path string
}

func (wh *WebHost) Frame(element, source string) string {
	var sb strings.Builder
	sb.WriteString("<")
	sb.WriteString(element)
	sb.WriteString(">\n")
	sb.WriteString(source)
	sb.WriteString("\n</")
	sb.WriteString(element)
	sb.WriteString(">\n")
	sb.WriteString("<script src=\"")
	sb.WriteString(wh.Base + wh.Tag + wh.Path) // full CDN path
	sb.WriteString(element)
	sb.WriteString(".js\"></script>\n")
	return sb.String()
}

type Record struct {
	TagName      string
	WebHost      WebHost
	Frame        interface{}
	RegisteredBy std.Realm
	RegisteredAt time.Time
	UpdatedAt    time.Time
}
