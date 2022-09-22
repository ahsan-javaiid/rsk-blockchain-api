package main

import (
    "log"
    "net/http"
	api "github.com/ahsan-javaiid/rsk-blockchain-api/api"
	config "github.com/ahsan-javaiid/rsk-blockchain-api/config"
)

func main() {
    config.LoadConfiguration("config/rpc.json")

    http.HandleFunc("/", api.Router)

    log.Println("Listening...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}