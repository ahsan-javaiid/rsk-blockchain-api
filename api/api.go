package api

import (
    "fmt"
	"io"
	"bytes"
	"log"
    "net/http"
	"encoding/json"
	"strings"
	config "github.com/ahsan-javaiid/rsk-blockchain-api/config"
)

var networkConfig map[string]string = map[string]string{
    "testnet": "https://public-node.testnet.rsk.co",
	"mainnet": "https://public-node.rsk.co",
}


func Router(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
		case "GET":
			Next(w, r)
		default:
			fmt.Fprintf(w, "Sorry, only GET method is supported")
    }
}

type ErrorResp struct {
	Status  string   `json:"status"`
	Msg     string   `json:"msg"`
}

func Next(w http.ResponseWriter, r *http.Request) {
    network, rpcName := splitLink(r.URL.Path, "/")
	rpcConfig := getRPCPayload(rpcName)

	if rpcConfig.Path == "" {
		w.Header().Set("Content-Type", "application/json")

		errResp := ErrorResp{ Status: "error", Msg: "rpc method not found"}
		
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResp)
		return
	}
	
	rpcJSON, err := json.Marshal(rpcConfig.Data)
	if err != nil {
		log.Fatal(err)
	}

	channel := make(chan http.Response)

	// Perform http request in go routine to support parallelism
    go httpRequest(networkConfig[network], channel, &rpcJSON)

	// Read response
	resp := <- channel 

	// Set response headers
	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

	// Send response
    io.Copy(w, resp.Body)

	// Close channel
	defer close(channel) 
	// Close http response body
	defer resp.Body.Close() 
}

func splitLink(s, sep string) (network string, rpcName string) {
    segments := strings.Split(s, sep)
	network = segments[1]
	rpcName = segments[2]
    
	return network, rpcName
}

func getRPCPayload(rpcName string) (config.RPC) {
	for _, element := range config.RpcList {
		if rpcName == element.Path {
			return element
		}
	}

	return config.RPC{}
}

func httpRequest(url string, c chan http.Response, payload *[]byte) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(*payload))

	if err != nil {
		log.Fatal(err)
	}

	c <- *resp
}