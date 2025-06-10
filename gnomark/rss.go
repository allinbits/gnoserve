package gnomark

import (
	"encoding/json"
	"strings"
)

func init() {
	RegisterTemplate("rss", RenderRss)
}

func getRssFormat(content string) string {
	_ = content // Use content if needed for processing
	return "rss"
}

func RenderRss(content string) string {
	format := getRssFormat(content)
	switch format {
	case "rss":
		return renderXmlRss(content)
	default:
		return "<p>Unsupported format: " + format + "</p>"
	}
}

type RssItem struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Content     string `json:"content:encoded"`
	Guid        string `json:"guid"`
	PubDate     string `json:"pubDate"`
}

func urlToAtomicLink(url string) string {
	linkType := "text/html"
	linkRel := "alternate"
	if strings.HasSuffix("feed", url) {
		linkType = "application/rss+xml"
		linkRel = "self"
	}
	if strings.HasSuffix("svg", url) {
		linkType = "image/svg+xml"
		linkRel = "alternate"
	}
	if strings.HasSuffix("json", url) {
		linkType = "application/json"
		linkRel = "alternate"
	}
	return `<atom:link href="` + url + `" rel="` + linkRel + `" type="` + linkType + `"/>`

}

func (i RssItem) Render(format string) string {
	if format == "rss" {
		return `<item>
	<title>` + i.Title + `</title>
	` + urlToAtomicLink(i.Link) + `
	<description>` + i.Description + `</description>
	<content:encoded>` + i.Content + `</content:encoded>
	<pubDate>` + i.PubDate + `</pubDate>
	<guid>` + i.Guid + `</guid>
	</item>`
	}
	return `<p>Unsupported format for item rendering: ` + format + `</p>`
}

func getItems(content string) []RssItem {
	var feed struct {
		GnoMark string `json:"gnoMark"`
		Format  string `json:"format"`
		Feed    struct {
			Title       string `json:"title"`
			Link        string `json:"link"`
			Description string `json:"description"`
			Created     string `json:"created"`
			Items       []struct {
				Title       string `json:"title"`
				Link        string `json:"link"`
				Description string `json:"description"`
				Content     string `json:"content"`
				Guid        string `json:"guid"`
				PubDate     string `json:"pubData"`
			} `json:"items"`
		} `json:"feed"`
	}

	err := json.Unmarshal([]byte(content), &feed)
	if err != nil {
		panic("Failed to parse content: " + err.Error())
	}

	var rssItems []RssItem
	for _, item := range feed.Feed.Items {
		rssItems = append(rssItems, RssItem{
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			Content:     item.Content,
			Guid:        item.Guid,
			PubDate:     item.PubDate,
		})
	}

	return rssItems
}

func getItemsXml(content string) string {
	items := getItems(content)
	var xmlItems string
	for _, item := range items {
		xmlItems += item.Render("rss")
	}
	return xmlItems
}

func renderXmlRss(content string) string {

	return `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/">
<channel>
  <title>GnoMark RSS Feed</title>
  <link>https://example.com/rss</link>
  <description>This is an example RSS feed for GnoMark.</description>
` + getItemsXml(content) + `
</channel>
</rss>`

}
