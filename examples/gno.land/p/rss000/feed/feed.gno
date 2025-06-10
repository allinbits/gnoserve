package feed

import (
	"strings"
	"time"
)

// Feed represents the main feed structure.
type Feed struct {
	Title       string
	Link        *Link
	Description string
	Author      *Author
	PubDate     time.Time
}

// Link represents a hyperlink with an Href attribute.
type Link struct {
	Href string
}

// Author represents the author of the feed or an item.
type Author struct {
	Name  string
	Email string
}

// Item represents an individual feed item.
type Item struct {
	Title       string
	Link        *Link
	Description string
	Content     string // Optional content for the item
	PubDate     time.Time
}

func (item *Item) ToJson() string {
	return `{
    "title": "` + item.Title + `",
    "link": "` + item.Link.Href + `",
    "description": "` + item.Description + `",
    "content": "` + item.Content + `",
    "guid": "` + item.Link.Href + `",
    "pubData": "` + item.PubDate.Format(time.RFC1123) + `"
}`
}

func itemsToJson(items []*Item) string {
	var jsonItems []string
	for _, item := range items {
		jsonItems = append(jsonItems, item.ToJson())
	}
	return strings.Join(jsonItems, ", ")
}

func (f *Feed) ToJson(items []*Item) string {
	return `{
        "title": "` + f.Title + `",
        "link": "` + f.Link.Href + `",
        "description": "` + f.Description + `",
        "created": "` + f.PubDate.Format(time.RFC3339) + `",
        "items": [` + itemsToJson(items) + `]
    }`
}
