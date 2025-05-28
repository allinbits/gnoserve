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
	PubDate:     time.Now(),
}

func frame(format string, data string) string {
	return "```gnomark\n\n" +
		`{ "gnoMark": "rss", "format": "` + format + `", "feed": ` + data + `}` +
		"\n```"
}

func getItems() (items []*feed.Item) {
	bmp.IterateRegistry(func(key string, value interface{}) bool {
		if item, ok := value.(bmp.ImageApi); ok {
			items = append(items, &feed.Item{
				Title:       key,
				Link:        &feed.Link{Href: "https://gnoserve.fly.dev/r/stackdump000/bmp:" + key},
				Description: item.Description(),
				Content:     item.Description(), // Optional content, can be empty
				PubDate:     time.Now(),
			})
		} else {
			panic("registry value is not an Item: " + key)
		}
		return false // continue iterating
	}, bmp.HasImageApi)
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
