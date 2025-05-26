package feed

import (
	"time"

	"gno.land/p/rss000/feed"
	"gno.land/r/stackdump000/bmp"
)

var rss = &feed.Feed{
	Title:       "Stack Dump",
	Link:        &feed.Link{Href: "https://gno.land/r/stackdump000"},
	Description: "A collection of stack dumps from Gno.land.",
	Author:      &feed.Author{Name: "Gno.land", Email: ""},
	Created:     time.Now(),
}

func frame(format string, data string) string {
	return "```gnomark\n\n" +
		`{ "gnoMark": "rss", "format": "` + format + `", "feed": ` + data + `}` +
		"\n```"
}

func getItems() (items []*feed.Item) {
    bmp.IterateRegistry(func(key string, value interface{}) bool {
        if item, ok := value.(*feed.Item); ok {
            items = append(items, item)
        }
        return true
    }, func(value interface{}) bool {
        _, ok := value.(*feed.Item)
        return ok
    })
    return items
}

func Render(path string) string {
	if path == "" {
		path = "md"
	}
	switch path {
	case "md":
	case "json":
	case "xml":
	case "html":
	case "htmlFragment":
	default:
		panic("unknown path")
	}
	items := getItems()
	return frame(path, rss.ToJson(items))
}
