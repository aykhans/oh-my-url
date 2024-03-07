package main

import (
	"github.com/aykhans/oh-my-url/app/config"
	"github.com/aykhans/oh-my-url/app/db"
	"github.com/aykhans/oh-my-url/app/http_handlers"
	"net/http"
	"sync"
)

func main() {
	config := config.GetAppConfig()
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
		panic(http.ListenAndServe(":"+config.LISTEN_PORT_CREATE, urlCreateMux))
	}()
	go func() {
		defer wg.Done()
		panic(http.ListenAndServe(":"+config.LISTEN_PORT_FORWARD, urlReadMux))
	}()
	wg.Wait()
}
