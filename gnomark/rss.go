package gnomark

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
	Title       string
	Link        string
	Description string
	PubDate     string
	Guid        string
}

func (i RssItem) Render(format string) string {
	if format == "rss" {
		return `<item>
	<title>` + i.Title + `</title>
	<link>` + i.Link + `</link>
	<description>` + i.Description + `</description>
	<pubDate>` + i.PubDate + `</pubDate>
	<guid>` + i.Guid + `</guid>
	</item>`
	}
	return `<p>Unsupported format for item rendering: ` + format + `</p>`
}

func getItems(content string) []RssItem {

	return []RssItem{
		{
			Title:       "Example Item",
			Link:        "https://example.com/item1",
			Description: "This is an example item in the GnoMark RSS feed.",
			PubDate:     "Mon, 01 Jan 2024 00:00:00 GMT",
			Guid:        "https://example.com/item1",
		},
	}
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
<rss version="2.0">
<channel>
  <title>GnoMark RSS Feed</title>
  <link>https://example.com/rss</link>
  <description>This is an example RSS feed for GnoMark.</description>
` + getItemsXml(content) + `
</channel>
</rss>`

}
