package feed

import (
"time"
"strings"
)


// Feed represents the main feed structure.
type Feed struct {
	Title       string
	Link        *Link
	Description string
	Author      *Author
	Created     time.Time
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
	Author      *Author
	Created     time.Time
}

func (item *Item) ToJson() string {
    return `{
        "title": "` + item.Title + `",
        "link": "` + item.Link.Href + `",
        "description": "` + item.Description + `",
        "author": {
            "name": "` + item.Author.Name + `",
            "email": "` + item.Author.Email + `"
        },
        "created": "` + item.Created.Format(time.RFC3339) + `"
    }`
}

func itemsToJson(items []*Item) string {
    var jsonItems []string
    for _, item := range items {
        jsonItems = append(jsonItems, item.ToJson())
    }
    return strings.Join(jsonItems, ", ")
}

func (f *Feed) ToJson(item []*Item) string {
    return `{
        "title": "` + f.Title + `",
        "link": "` + f.Link.Href + `",
        "description": "` + f.Description + `",
        "author": {
            "name": "` + f.Author.Name + `",
            "email": "` + f.Author.Email + `"
        },
        "created": "` + f.Created.Format(time.RFC3339) + `",
        "items": [` + itemsToJson(item) + `]
    }`
}