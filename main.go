package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:1337", nil)
}

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, _ := getJsonResponse()
	fmt.Fprintf(w, string(response))
}

func getJsonResponse() ([]byte, error) {
	d := make(map[string]string)
	d["hello"] = "world"
	return json.Marshal(d)
}
