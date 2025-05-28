package feed

import (
	"time"

	"gno.land/p/rss000/feed"
	"gno.land/r/stackdump000/bmp"
)

var rss = &feed.Feed{
	Title:       "GnoCap: bitmaps on Gno.land",
	Link:        &feed.Link{Href: "https://gnoserve.fly.dev/rss/r/stackdump000/feed"},
	Description: "A collection of bitmaps stored on Gno.land.",
	Author:      &feed.Author{Name: "stackdump", Email: "myork@stackdump.com"},
	Created:    "2025-05-27T00:00:00Z",
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
