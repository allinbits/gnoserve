package gnomark

import (
	"std"
	"time"
)

type Frame interface {
    ToGnoMark(string) string
}

type WebHost struct {
	Base string
	Tag  string
	Path string
}

func (wh *WebHost) Frame(element, source string) string {
   return "```gnomark\n" + source + "\n```\n"
}

type Record struct {
	TagName      string
	WebHost      WebHost
	RegisteredBy std.Realm
	RegisteredAt time.Time
	UpdatedAt    time.Time
}
