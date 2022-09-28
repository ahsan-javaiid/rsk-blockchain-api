package main

import (
    "log"
    "fmt"
    "os"
    "net/http"
	api "github.com/ahsan-javaiid/rsk-blockchain-api/api"
	config "github.com/ahsan-javaiid/rsk-blockchain-api/config"
    utils "github.com/ahsan-javaiid/rsk-blockchain-api/utils"
)

func main() {
    config.LoadConfiguration("config/rpc.json")

    http.HandleFunc("/", api.Router)

    if !utils.IsEnvExist("PORT") {
        os.Setenv("PORT", "8080")
    }

    PORT := os.Getenv("PORT")

    log.Println("Listening on port ",PORT)

    log.Fatal(http.ListenAndServe(fmt.Sprint(":", PORT), nil))
}