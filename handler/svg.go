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

func (h *WebHandler) RenderSvg(w http.ResponseWriter, r *http.Request) {
	// Use a buffer to render the realm
	var content bytes.Buffer
	gnourl, _ := ParseSvgUrl(r)
	_, err := h.Client.RenderRealm(&content, gnourl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write(content.Bytes())
}
