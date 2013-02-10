package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// documentation for csv is at http://golang.org/pkg/encoding/csv/
// TODO: could not find
func main() {
	file, error := os.Open("data/sample.csv")
	if error != nil {
		// error is printable
		// elements passed are separated by space automatically
		fmt.Println("Error:", error)
		return
	}
	// automatically call Close() at the end of current method
	defer file.Close()
	// 
	reader := csv.NewReader(file)
	// options are available at:
	// http://golang.org/src/pkg/encoding/csv/reader.go?s=3213:3671#L94
	reader.Comma = ';'
	lineCount := 0
	for {
		// read just one record, but we could ReadAll() as well
		record, error := reader.Read()
		// end-of-file is fitted into error
		if error == io.EOF {
			break
		} else if error != nil {
			fmt.Println("Error:", error)
			return
		}
		// record is an array of string so is directly printable
		fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
		// and we can iterate on top of that
		for i := 0; i < len(record); i++ {
			fmt.Println(" ", record[i])
		}
		fmt.Println()
		lineCount += 1
	}
}
