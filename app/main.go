package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/aykhans/oh-my-url/app/config"
	"github.com/aykhans/oh-my-url/app/db"
	"github.com/aykhans/oh-my-url/app/http_handlers"
)

func main() {
	appConfig := config.GetAppConfig()
	dbCreate := db.GetDB()
	dbCreate.Init()
	handlerCreate := httpHandlers.HandlerCreate{DB: dbCreate}
	urlCreateMux := http.NewServeMux()
	urlCreateMux.HandleFunc("/", handlerCreate.UrlCreate)
	urlCreateMux.HandleFunc("/favicon.ico", httpHandlers.FaviconHandler)

	dbRead := db.GetDB()
	handlerForward := httpHandlers.HandlerForward{DB: dbRead}
	urlReadMux := http.NewServeMux()
	urlReadMux.HandleFunc("/", handlerForward.UrlForward)
	urlReadMux.HandleFunc("/favicon.ico", httpHandlers.FaviconHandler)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			err := http.ListenAndServe(":"+appConfig.LISTEN_PORT_CREATE, urlCreateMux)
			if err != nil {
				log.Println(err)
				continue
			}
			break
		}
	}()
	go func() {
		defer wg.Done()
		for {
			err := http.ListenAndServe(":"+appConfig.LISTEN_PORT_FORWARD, urlReadMux)
			if err != nil {
				log.Println(err)
				continue
			}
			break
		}
	}()
	wg.Wait()
}
