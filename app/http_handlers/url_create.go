package httpHandlers

import (
	"html/template"
	"log"
	"net/http"
	netUrl "net/url"
	"regexp"

	"github.com/aykhans/oh-my-url/app/config"
	"github.com/aykhans/oh-my-url/app/errors"
	"github.com/aykhans/oh-my-url/app/utils"
)

type CreateData struct {
	ShortedURL string
	MainURL    string
	Error      error
}

func (hl *HandlerCreate) UrlCreate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(utils.GetTemplatePaths("index.html")...)
	if err != nil {
		log.Println(err)
		InternalServerError(w, errors.ErrAPITemplateParsing)
		return
	}

	switch r.Method {
	case http.MethodGet:
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			InternalServerError(w, errors.ErrAPITemplateParsing)
			return
		}
	case http.MethodPost:
		url := r.FormValue("url")
		urlRegex := regexp.MustCompile(`^(http|https)://[a-zA-Z0-9.-]+(?:\:[0-9]+)?(?:/[^\s]*)?$`)
		isValidUrl := urlRegex.MatchString(url)
		if !isValidUrl {
			data := CreateData{
				MainURL: url,
				Error:   errors.ErrAPIInvalidURL,
			}
			err = tmpl.Execute(w, data)
			if err != nil {
				log.Println(err)
				InternalServerError(w, errors.ErrAPITemplateParsing)
			}
			return
		}

		key, err := hl.DB.CreateURL(url)
		if err != nil {
			data := CreateData{
				MainURL: url,
				Error:   errors.ErrAPI503,
			}
			err = tmpl.Execute(w, data)
			if err != nil {
				log.Println(err)
				InternalServerError(w, errors.ErrAPITemplateParsing)
			}
			return
		}
		shortedURL, err := netUrl.JoinPath(config.GetForwardDomain(), key)
		if err != nil {
			log.Println(err)
			InternalServerError(w, errors.ErrAPITemplateParsing)
			return
		}
		data := CreateData{
			ShortedURL: shortedURL,
			MainURL:    url,
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			InternalServerError(w, errors.ErrAPITemplateParsing)
			return
		}

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
