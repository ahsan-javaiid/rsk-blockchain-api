package api

import (
    "fmt"
	"bytes"
	"log"
    "net/http"
	"encoding/json"
	"strings"
	config "github.com/ahsan-javaiid/rsk-blockchain-api/config"
)

var network map[string]string = map[string]string{
    "testnet": "https://public-node.testnet.rsk.co",
	"mainnet": "https://public-node.rsk.co	",
}


func Router(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rpc json:", config.RpcList)

    // if r.URL.Path != "/" {
    //     http.Error(w, "404 not found.", http.StatusNotFound)
    //     return
    // }

    switch r.Method {
    
		case "GET":
			Next(w, r)
		default:
			fmt.Fprintf(w, "Sorry, only GET method is supported")
    }
}

// http://localhost:8080/testnet/web3_clientVersion
func Next(w http.ResponseWriter, r *http.Request) {
    network, rpcName := splitLink(r.URL.Path, "/")
	rpcConfig := getRPCPayload(rpcName)
	rpcJSON, err := json.Marshal(rpcConfig.Data)
	if err != nil {
		log.Fatal(err)
	}
    // Todo: send post request
	// https://zetcode.com/golang/getpostrequest/
	fmt.Println(network, rpcConfig, rpcJSON)
	fmt.Println("---------")
	fmt.Println("---------", bytes.NewBuffer(rpcJSON))
	w.Header().Set("Content-Type", "application/json")
	var data = make(map[string]string)
	data["msg"] = "done"
	json.NewEncoder(w).Encode(data)
}

func splitLink(s, sep string) (network string, rpcName string) {
    segments := strings.Split(s, sep)
	network = segments[1]
	rpcName = segments[2]
    return network, rpcName
}

func getRPCPayload(rpcName string) (config.RPC) {
	for index, element := range config.RpcList {
		fmt.Println(index, element)
		if rpcName == element.Path {
			return element
		}
	}

	return config.RPC{}
}