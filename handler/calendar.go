package handler

import "net/http"

// get the target realm from the request and use it to get the
// target realm from the client and then template the index.gno event in an ical format
func (h *WebHandler) RenderCalendar(w http.ResponseWriter, r *http.Request) {
	// get the target realm from the request
	realm := r.URL.Query().Get("realm")
	if realm == "" {
		http.Error(w, "missing realm", http.StatusBadRequest)
		return
	}
	// FIXME get the target realm from the client
	// gno call Calendar(path) and write output
}
