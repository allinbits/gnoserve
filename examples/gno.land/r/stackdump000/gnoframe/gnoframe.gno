package gnoframe

import (
    frame "gno.land/r/frames000"
)

var feed = RssFeed{}

type RssFeed struct {}

func (RssFeed) ToGnoMark(_ string) string {
    return rssData // Hardcoded RSS data for demonstration purposes
}

var rssData = `
    { "gnoMark": "rss", "format": "md", "feed": {
    "title": "GnoCap: bitmaps on Gno.land",
    "link": "https://gnoserve.fly.dev/rss/r/stackdump000/feed",
    "description": "A collection of bitmaps stored on Gno.land.",
    "created": "10 Jun 2025 11:10 +0000",
    "items": [{
    "title": "aib",
    "link": "https://gnoserve.fly.dev/r/stackdump000/bmp:aib",
    "description": "Aib | AIB logo | 50x50 pixels",
    "content": "",
    "guid": "https://gnoserve.fly.dev/r/stackdump000/bmp:aib",
    "pubData": "10 Jun 2025 11:11 +0000"
    }, {
    "title": "gnome",
    "link": "https://gnoserve.fly.dev/r/stackdump000/bmp:gnome",
    "description": "Gnome | A pixelated gnome",
    "content": "",
    "guid": "https://gnoserve.fly.dev/r/stackdump000/bmp:gnome",
    "pubData": "10 Jun 2025 11:11 +0000"
    }, {
    "title": "profile",
    "link": "https://gnoserve.fly.dev/r/stackdump000/bmp:profile",
    "description": "GnoCap: a creative commons | collectively editing 250x250 pixels",
    "content": "",
    "guid": "https://gnoserve.fly.dev/r/stackdump000/bmp:profile",
    "pubData": "10 Jun 2025 11:11 +0000"
    }, {
    "title": "sierpinski",
    "link": "https://gnoserve.fly.dev/r/stackdump000/bmp:sierpinski",
    "description": "Sierpinski's triangle | Recursive fractal pattern - an algorithmic pixel set",
    "content": "",
    "guid": "https://gnoserve.fly.dev/r/stackdump000/bmp:sierpinski",
    "pubData": "10 Jun 2025 11:11 +0000"
    }]
    }}
`

func Render(path string) string {
    return frame.RenderTag("rss-feed", path, feed)
}