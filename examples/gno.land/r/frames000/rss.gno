package frame

import (
	"strings"
)

var rssItems []RssItem

func init() {
	rssItems = []RssItem{
		{
			Title:       "Example Item",
			Link:        "https://example.com/item1",
			Description: "This is an example item description.",
			Content:     "<p>This is the content of the example item.</p>",
			Guid:        "item1-guid",
			PubDate:     "Mon, 01 Jan 2024 00:00:00 GMT",
		},
	}

	register("rss-feed", webHost, RssFrame)
}

type RssItem struct {
	Title       string
	Link        string
	Description string
	Content     string
	Guid        string
	PubDate     string
}

func RssFrame(path string, props ...interface{}) string {
	if len(props) > 0 {
		if items, ok := props[0].([]RssItem); ok {
			rssItems = items
		}
	}

	return renderXmlRss(rssItems)
}

func urlToAtomicLink(url string) string {
	linkType := "text/html"
	linkRel := "alternate"

	if strings.HasSuffix(url, "feed") {
		linkType = "application/rss+xml"
		linkRel = "self"
	}
	if strings.HasSuffix(url, "svg") {
		linkType = "image/svg+xml"
		linkRel = "alternate"
	}
	if strings.HasSuffix(url, "json") {
		linkType = "application/json"
		linkRel = "alternate"
	}
	return `<atom:link href="` + url + `" rel="` + linkRel + `" type="` + linkType + `"/>`
}

func (i RssItem) Render(_ string) string {
	return `<item>
    <title>` + i.Title + `</title>
    ` + urlToAtomicLink(i.Link) + `
    <description>` + i.Description + `</description>
    <content:encoded>` + i.Content + `</content:encoded>
    <pubDate>` + i.PubDate + `</pubDate>
    <guid>` + i.Guid + `</guid>
</item>`
}

func getItemsXml(items []RssItem) string {
	var xmlItems string
	for _, item := range items {
		xmlItems += item.Render("rss")
	}
	return xmlItems
}

func renderXmlRss(items []RssItem) string {
	feedUrl := "https://gnoserve.fly.dev/rss/r/stackdump000/feed"

	return `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom">
<channel>
  <title>GnoMark RSS Feed</title>
  <link>` + feedUrl + `</link>
` + urlToAtomicLink(feedUrl) + `
  <description>This is an example RSS feed for GnoMark.</description>
` + getItemsXml(items) + `
</channel>
</rss>`
}
