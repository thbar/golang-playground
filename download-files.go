package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadFromUrl(url string) {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Println("Downloading", url, "to", fileName)

	// TODO: check file existence first with io.IsExist
	output, error := os.Create(fileName)
	if error != nil {
		fmt.Println("Error while creating", fileName, "-", error)
		return
	}
	defer output.Close()

	response, error := http.Get(url)
	if error != nil {
		fmt.Println("Error while downloading", url, "-", error)
		return
	}
	defer response.Body.Close()

	n, error := io.Copy(output, response.Body)
	if error != nil {
		fmt.Println("Error while downloading", url, "-", error)
		return
	}

	fmt.Println(n, "bytes downloaded.")
}

func main() {
	countries := []string{"GB", "FR", "ES", "DE", "CN", "CA", "ID", "US"}
	for i := 0; i < len(countries); i++ {
		url := "http://download.geonames.org/export/dump/" + countries[i] + ".zip"
		downloadFromUrl(url)
	}
}
