package config

import (
	"encoding/json"
	"log"
	"os"
	"io/ioutil"
)



type RPC struct {
    Data struct {
        Id          int      `json:"id"`
        JsonRPC     string   `json:"jsonrpc"`
        Method      string   `json:"method"`
        Params      []string `json:"params"`
    }                        `json:"data"`
    Path            string   `json:"path"`
}

var RpcList []RPC

func LoadConfiguration(file string) []RPC {
	filename, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	defer filename.Close()

	data, err := ioutil.ReadAll(filename)

	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal(data, &RpcList)

	if jsonErr != nil {
        log.Fatal(jsonErr)
    }

	return RpcList
}