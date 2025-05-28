package handler

import (
	"bytes"
	"github.com/gnolang/gno/gno.land/pkg/gnoweb/weburl"
	"net/http"
)

func ParseRssUrl(r *http.Request) (*weburl.GnoURL, error) {
	gnourl, _ := weburl.ParseFromURL(r.URL)
	// remove leading /rss/ from gnourl.Path
	gnourl.Path = gnourl.Path[4:]
	return gnourl, nil
}

func (h *WebHandler) RenderRss(w http.ResponseWriter, r *http.Request) {
	// Use a buffer to render the realm
	var content bytes.Buffer
	gnourl, _ := ParseRssUrl(r)
	_, err := h.Client.RenderRealm(&content, gnourl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")
	w.Header().Set("X-Realm-Path", gnourl.Path)
	w.Header().Set("X-Realm-Args", gnourl.Args)
	w.Write(content.Bytes())
}
