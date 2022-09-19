package api

import (
    "fmt"
    "net/http"
	"encoding/json"
)

func Process(w http.ResponseWriter, r *http.Request) {

    if r.URL.Path != "/" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    switch r.Method {
    
		case "GET":
			Next(w, r)
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}

func Next(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var data = make(map[string]string)
	data["msg"] = "done"
	json.NewEncoder(w).Encode(data)
}