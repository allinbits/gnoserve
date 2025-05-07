package www

import "net/url"

var (
	jsonLd = `{
      "@context": "https://schema.org"
}`

	viewSourceSVGButton = `<svg xmlns="http://www.w3.org/2000/svg" width="200" height="34" viewBox="0 0 200 24">
<rect width="110" height="24" fill="#f0f0f0" stroke="#cccccc" stroke-width="1" rx="4" ry="4" filter="url(#shadow)"/>
<text x="10" y="18" font-family="Arial, sans-serif" font-size="16" fill="#000000">View Source</text>
<defs>
  <filter id="shadow" x="-10%" y="-10%" width="120%" height="120%">
    <feOffset result="offOut" in="SourceAlpha" dx="2" dy="2" />
    <feGaussianBlur result="blurOut" in="offOut" stdDeviation="2" />
    <feBlend in="SourceGraphic" in2="blurOut" mode="normal" />
  </filter>
</defs>
</svg>`

	dataUrl = "data:image/svg+xml;utf-8," + url.PathEscape(viewSourceSVGButton)

	message = "## jsonld Example\n\n" +
		"This is a simple example of JSON-LD formatted data.\n\n" +
		"which is inserted into the page using &lt;script type=\"application/ld+json\"&gt;.\n\n" +
		"view the source to see how this is done:" + "[![View Source](" + dataUrl + ")](jsonld$source)"
)

func Render(_ string) string {
	return message + "\n\n" +
		"```jsonld\n\n" + jsonLd + "\n\n```"
}
