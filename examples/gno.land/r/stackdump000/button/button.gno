package button

import (
	"net/url"
	"strconv"
	"strings"
)

func Render(path string) string {
	// Example usage of the Button function
	return SubmitButton("Click Me really big button", path, 16, 200) +
		SubmitButton("Submit a Information, Media & Social Infrastructure Proposal", path, 16, 200) +
		SubmitButton("Submit a Information, Media & Social Infrastructure Proposal", path, 26, 200)
}

func escapeHtml(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	s = strings.ReplaceAll(s, "'", "&#39;")
	return s
}

func SubmitButton(label, path string, fontSize, minWidth int) string {
	charWidth := int(0.6 * float64(fontSize)) // Approximate width of each character
	padding := 40                             // Total padding (left + right)
	h := 2 * fontSize                         // Height of the button, 2x font size for padding
	w := len(label)*charWidth + padding
	if w < minWidth {
		w = minWidth
	}

	svgButton := `<svg xmlns="http://www.w3.org/2000/svg" width="` + strconv.Itoa(w) + `" height="` + strconv.Itoa(h) + `">
<defs>
  <filter id="dropShadow" x="-50%" y="-50%" width="200%" height="200%">
    <feGaussianBlur in="SourceAlpha" stdDeviation="3" result="blur"/>
    <feOffset in="blur" dx="2" dy="2" result="offsetBlur"/>
    <feMerge>
      <feMergeNode in="offsetBlur"/>
      <feMergeNode in="SourceGraphic"/>
    </feMerge>
  </filter>
</defs>
<foreignObject x="16" y="-5" width="` + strconv.Itoa(w) + `" height="` + strconv.Itoa(h) + `" filter="url(#dropShadow)">
  <body xmlns="http://www.w3.org/1999/xhtml">
    <button style="padding-left: 20px; font-size:` + strconv.Itoa(fontSize) + `px">
      ` + escapeHtml(label) + `
    </button>
  </body>
</foreignObject>
</svg>`

	dataUrl := "data:image/svg+xml;utf8," + url.PathEscape(svgButton)

	return "[![" + label + "](" + dataUrl + ")](" + path + ")"
}
