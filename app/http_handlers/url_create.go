package httpHandlers

import (
	"github.com/aykhans/oh-my-url/app/utils"
	"github.com/aykhans/oh-my-url/app/config"
	"html/template"
	"net/http"
	netUrl "net/url"
	"regexp"
)

type CreateData struct {
	ShortedURL string
	MainURL    string
	Error      string
}

func (hl *HandlerCreate) UrlCreate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(utils.GetTemplatePaths("index.html")...)
	if err != nil {
		InternalServerError(w, err)
		return
	}

	switch r.Method {
	case http.MethodGet:
		err = tmpl.Execute(w, nil)
		if err != nil {
			InternalServerError(w, err)
			return
		}
	case http.MethodPost:
		url := r.FormValue("url")
		urlRegex := regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+(?:\:[0-9]+)?(?:/[^\s]*)?$`)
		isValidUrl := urlRegex.MatchString(url)
		if !isValidUrl {
			data := CreateData{
				MainURL: url,
				Error:   "Invalid URL",
			}
			err = tmpl.Execute(w, data)
			if err != nil {
				InternalServerError(w, err)
			}
			return
		}

		key, err := hl.DB.CreateURL(url)
		if err != nil {
			InternalServerError(w, err)
			return
		}
		shortedURL, err := netUrl.JoinPath(config.GetForwardDomain(), key)
		if err != nil {
			InternalServerError(w, err)
			return
		}
		data := CreateData{
			ShortedURL: shortedURL,
			MainURL:    url,
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			InternalServerError(w, err)
			return
		}

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
