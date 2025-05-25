package handler

import (
	"bytes"
	"github.com/gnolang/gno/gno.land/pkg/gnoweb/weburl"
	"net/http"
)

func ParseSvgUrl(r *http.Request) (*weburl.GnoURL, error) {
	gnourl, _ := weburl.ParseFromURL(r.URL)
	// remove leading /svg/ from gnourl.Path
	gnourl.Path = gnourl.Path[4:]
	return gnourl, nil
}

func IsValidSvg(content bytes.Buffer) bool {
	reader := bytes.NewReader(content.Bytes())
	for {
		b, err := reader.ReadByte()
		if err != nil {
			return false // EOF or error
		}
		if !bytes.Contains([]byte{b}, []byte{' ', '\n', '\t', '\r'}) {
			// Check if the first non-whitespace character starts with "<svg"
			if b == '<' {
				next := make([]byte, 3)
				if _, err := reader.Read(next); err != nil {
					return false
				}
				return string(next) == "svg"
			}
			return false
		}
	}
}

func (h *WebHandler) RenderSvg(w http.ResponseWriter, r *http.Request) {
	// Use a buffer to render the realm
	var content bytes.Buffer
	gnourl, _ := ParseSvgUrl(r)
	_, err := h.Client.RenderRealm(&content, gnourl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// if content doesn't start with <?xml return an error
	if !IsValidSvg(content) {
		http.Error(w, "Invalid SVG content", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("X-Realm-Path", gnourl.Path)
	w.Header().Set("X-Realm-Args", gnourl.Args)
	w.Write(content.Bytes())
}
