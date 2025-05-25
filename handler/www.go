package handler

import (
	"bytes"
	"net/http"
)

func (h *WebHandler) RenderHtml(w http.ResponseWriter, r *http.Request) {
	// Use a buffer to render the realm
	var content bytes.Buffer
	gnourl, _ := ParseSvgUrl(r) // Reuse ParseSvgUrl to parse the URL
	_, err := h.Client.RenderRealm(&content, gnourl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Realm-Path", gnourl.Path)
	w.Header().Set("X-Realm-Args", gnourl.Args)
	w.Write(content.Bytes())
}
