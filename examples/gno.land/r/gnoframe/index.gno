package gnoframe

var frame = `
    {
        "gnoMark": "frame",
        "frame": "gno-frame",
        "version": "0.1.0",
        "name": "GnoFrame",
        "iconUrl": "https://cdn.jsdelivr.net/gh/allinbits/gnoserve@0.1.0/static/favicon.ico",
        "homeUrl": "https://gno.land",
        "imageUrl": "https://cdn.jsdelivr.net/gh/allinbits/gnoserve@0.1.0/static/android-chrome-512x512.png",
        "buttonTitle": "GnoFrame",
        "splashImageUrl": "",
        "splashBackgroundColor": "#ffffff",
        "webhookUrl": "https://gno.land/webhook"
    }
`

var gnoMarkStart = "```gnomark"
var gnoMarkEnd = "```"

func Render(_ string) string {
	return "## GnoMark: frame\n\n" + gnoMarkStart + "\n\n" + frame + "\n\n" + gnoMarkEnd
}
