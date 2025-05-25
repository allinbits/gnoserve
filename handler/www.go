package handler

import (
	"bytes"
	"net/http"
)

// FIXME: http://127.0.0.1:8888/www/r/stackdump000/bmp:profile:html
// has an error on line 115 - note: though unsafeHTML is enabled this still
// passes throught the rendering engine - it may be replacing chars? or something else
// TODO: test the pixelizer app as stand-alone html
func IsValidHtml(content bytes.Buffer) bool {
	return true
	reader := bytes.NewReader(content.Bytes())
	for {
		b, err := reader.ReadByte()
		if err != nil {
			return false // EOF or error
		}
		if !bytes.Contains([]byte{b}, []byte{' ', '\n', '\t', '\r'}) {
			// Check if the first non-whitespace character starts with "<!DOCTYPE html"
			if b == '<' {
				next := make([]byte, 15)
				if _, err := reader.Read(next); err != nil {
					return false
				}
				return string(next) == "!DOCTYPE html"
			}
			return false
		}
	}
}

func (h *WebHandler) RenderHtml(w http.ResponseWriter, r *http.Request) {
	// Use a buffer to render the realm
	var content bytes.Buffer
	gnourl, _ := ParseSvgUrl(r) // Reuse ParseSvgUrl to parse the URL
	_, err := h.Client.RenderRealm(&content, gnourl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Validate if content starts with <!DOCTYPE html
	if !IsValidHtml(content) {
		http.Error(w, "Invalid HTML content", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Realm-Path", gnourl.Path)
	w.Header().Set("X-Realm-Args", gnourl.Args)
	w.Write(content.Bytes())
}
