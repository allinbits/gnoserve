package gnomark

import (
	"std"
	"time"
	"strings"
)

type Frame interface {
    ToGnoMark(string) string
}

type WebHost struct {
	Base string
	Tag  string
	Path string
}

func dataKey(tag string) string {
    parts := strings.Split(tag, "-")
    if len(parts) > 0 {
        return parts[1]
    }
    panic("unexpected tag format")
}

func (wh *WebHost) Frame(element, format, source string) string {
    jsonData := `{"gnoMark": "` + element + `", "format": "` + format + `", "` + dataKey(element) + `": ` + source + `}`
    return "```gnomark\n\n" + jsonData + "\n\n```"
}

type Record struct {
	TagName      string
	WebHost      WebHost
	Description  string
	RegisteredBy std.Realm
	RegisteredAt time.Time
	UpdatedAt    time.Time
}
