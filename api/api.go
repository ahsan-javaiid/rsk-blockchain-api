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

	resp, err := http.Post(networkConfig[network], "application/json", bytes.NewBuffer(rpcJSON))

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
    io.Copy(w, resp.Body)
    resp.Body.Close()
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