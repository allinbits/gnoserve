package feed

import "time"

// Feed represents the main feed structure.
type Feed struct {
	Title       string
	Link        *Link
	Description string
	Author      *Author
	Created     time.Time
	Items       []*Item
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
