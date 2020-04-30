package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error
}

var csvData []CSVRecord

func main() {
	f, err := os.Open("data/iris.data")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)

	reader.FieldsPerRecord = -1

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		var csvRecord CSVRecord

		for idx, value := range record {
			if idx == 4 {
				if value == "" {
					log.Printf("Unexpected type in column %d\n", idx)
					csvRecord.ParseError = fmt.Errorf("Empty string value")
					break
				}

				csvRecord.Species = value
				continue
			}

			var floatValue float64

			if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Unexpected type in column %d\n", idx)
				csvRecord.ParseError = fmt.Errorf("Could not parse float")
				break
			}

			switch idx {
			case 0:
				csvRecord.SepalLength = floatValue
			case 1:
				csvRecord.SepalWidth = floatValue
			case 2:
				csvRecord.PetalLength = floatValue
			case 3:
				csvRecord.PetalWidth = floatValue
			}

		}

		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}
	}

	fmt.Println(csvData)

}
