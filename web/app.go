package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// http://golang.org/pkg/net/http/

// serve static page for one file manually (but we could do better here)
func HtmlHandler(response http.ResponseWriter, req *http.Request) {
	var result, err = ioutil.ReadFile("index.html")
	if err != nil {
		// one can use http.Error to stop the processing and return an error
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "text/html")
	response.Header().Set("Content-Length", strconv.Itoa(len(result)))
	fmt.Fprintf(response, "%s", result)
}

// use a struct to define my JSON output
type Fruit struct {
	Name              string  `json:"name"`
	Price             float64 `json:"unit_price"`
	AvailableQuantity int     `json:"available_quantity"`
}

func JSONHandler(response http.ResponseWriter, req *http.Request) {
	var fruits = []Fruit{
		Fruit{"Orange", 5.47, 150},
		Fruit{"Apple", 1.20, 300},
		Fruit{"Banana", 2.54, 37},
		Fruit{"Kumquat", 3.57, 5}}
	// MarshalIndent let us pretty print json
	b, _ := json.MarshalIndent(fruits, "", " ") // TODO: handle error here
	response.Header().Set("Content-Type", "application/json")
	// TODO: set content length here
	response.Write(b)
}

func main() {
	port := "3000"
	http.HandleFunc("/", http.HandlerFunc(HtmlHandler))
	http.HandleFunc("/data.json", http.HandlerFunc(JSONHandler))
	fmt.Println("Listening on port", port)
	http.ListenAndServe((":" + port), nil)
}
