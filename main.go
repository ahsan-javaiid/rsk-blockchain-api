package main

import (
    "log"
    "net/http"
	api "github.com/ahsan-javaiid/rsk-blockchain-api/api"
)



func main() {

    http.HandleFunc("/", api.Process)

    log.Println("Listening...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}