package handler

import "net/http"

func (h *WebHandler) RenderRss(w http.ResponseWriter, r *http.Request) {
	_ = r // TODO: expand to read realm from request
	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Example RSS feed content
	rssContent := `<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0">
<channel>
  <title>GnoMark RSS Feed</title>
  <link>https://example.com/rss</link>
  <description>This is an example RSS feed for GnoMark.</description>
  <item>
	<title>Example Item</title>
	<link>https://example.com/item1</link>
	<description>This is an example item in the GnoMark RSS feed.</description>
	<pubDate>Mon, 01 Jan 2024 00:00:00 GMT</pubDate>
	<guid>https://example.com/item1</guid>
  </item>
  <item>
	<title>Another Example Item</title>
	<link>https://example.com/item2</link>
	<description>This is another example item in the GnoMark RSS feed.</description>
	<pubDate>Tue, 02 Jan 2024 00:00:00 GMT</pubDate>
	<guid>https://example.com/item2</guid>
  </item>
</channel>
</rss>`
	// Write the RSS content to the response
	_, err := w.Write([]byte(rssContent))
	if err != nil {
		http.Error(w, "Failed to write RSS content", http.StatusInternalServerError)
		return
	}
}
