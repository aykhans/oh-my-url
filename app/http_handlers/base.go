package httpHandlers

import (
	"github.com/aykhans/oh-my-url/app/db"
	"github.com/aykhans/oh-my-url/app/utils"
	"log"
	"net/http"
)

type HandlerCreate struct {
	DB db.DB
}

type HandlerForward struct {
	DB db.DB
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, utils.GetTemplatePaths("favicon.ico")[0])
}

func InternalServerError(w http.ResponseWriter, err error) {
	log.Fatal(err)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
