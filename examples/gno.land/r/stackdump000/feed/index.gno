package feed

import (
	"time"

	"gno.land/p/rss000/feed"
)

var rss = &feed.Feed{
	Title:       "Stack Dump",
	Link:        &feed.Link{Href: "https://gno.land/r/stackdump000"},
	Description: "A collection of stack dumps from Gno.land.",
	Author:      &feed.Author{Name: "Gno.land", Email: ""},
	Created:     time.Now(),
}

type Rss struct {
	*feed.Feed
}

func (Rss) Items() []*feed.Item {
	return rss.Items
}

func Render(path string) string {
    switch path {
    case "xml":
        return renderFeed()
    default:
        return "```gnomark\n\n{ \"gnoMark\": \"rss\" }\n\n```"
    }
}
