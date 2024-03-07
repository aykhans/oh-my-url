package httpHandlers

import (
	"github.com/aykhans/oh-my-url/app/config"
	"net/http"
	"strings"
)

func (hl *HandlerForward) UrlForward(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	segments := strings.Split(path, "/")
	if len(segments) > 2 {
		http.NotFound(w, r)
		return
	} else if segments[1] == "" {
		http.Redirect(w, r, config.GetCreateDomain(), http.StatusMovedPermanently)
		return
	}

	key := segments[1]
	url, err := hl.DB.GetURL(key)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
