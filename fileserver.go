package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
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
	http.Handle("/", http.FileServer(http.Dir(ServerCfg.ServerRoot)))
	http.ListenAndServe(":"+strconv.Itoa(ServerCfg.Port), nil)
}
