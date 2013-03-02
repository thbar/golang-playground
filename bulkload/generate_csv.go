package main

import "encoding/csv"
import "os"

import "flag"
import "strconv"

func main() {
	numberOfRecords := flag.Int("count", 15000000, "number of records to be generated")
	flag.Parse()

	file, err := os.Create("go-data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"guid", "some_field", "some_other_field"})
	for i := 0; i < *numberOfRecords; i++ {
		record := []string{
			"guid_" + strconv.Itoa(i),
			"some_field_" + strconv.Itoa(i),
			"some_other_field_" + strconv.Itoa(i)}
		writer.Write(record)
	}
}
