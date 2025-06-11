package handler

import (
	"bytes"
	"github.com/gnolang/gno/gno.land/pkg/gnoweb/weburl"
	"net/http"
	"net/url"
)

type GnoURL struct {
	// Example full path:
	// https://labsnet.fly.dev/r/buidlthefuture000/events/gnolandlaunch:edit?event=1&session=2

	Domain   string     // gno.land
	Path     string     // r/buidlthefuture000/events/gnolandlaunch
	Args     string     // edit
	WebQuery url.Values // help&a=b // what causes this??
	Query    url.Values // event=1&session=2
	File     string     // gnolandlaunch // has .gno??
}

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
	// REVIEW: appcfg.UnsafeHTML is required since w/ are calling RenderRealm
	_, err := h.Client.RenderRealm(&content, gnourl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("X-Realm-Path", gnourl.Path)
	w.Header().Set("X-Realm-Args", gnourl.Args)
	w.Write(content.Bytes())
}
