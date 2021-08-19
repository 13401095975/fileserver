package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	configName := flag.String("config", "server.json", "config file name")
	flag.Parse()

	err := LoadConfig(*configName)
	if err != nil {
		log.Fatal("Load config file Failed,", err)
		os.Exit(1)
	}

	log.Printf("%s server running at port: %d", ServerCfg.AppName, ServerCfg.Port)
	http.Handle("/", CacheControlWrapper(http.FileServer(http.Dir(ServerCfg.ServerRoot))))
	http.ListenAndServe(":"+strconv.Itoa(ServerCfg.Port), nil)
}

func CacheControlWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if len(ServerCfg.CacheFileType) == 0 {
			h.ServeHTTP(w, r)
			return
		}
		for _, suffix := range ServerCfg.CacheFileType {
			if strings.HasSuffix(r.RequestURI, suffix) {
				w.Header().Set("Cache-Control", "max-age=2592000")
				break
			}
		}

		h.ServeHTTP(w, r)
	})
}
