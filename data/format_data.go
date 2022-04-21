package data

import (
	"encoding/csv"
	"fmt"
	"os"
)

func CsvToArray(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	r.LazyQuotes = true
	r.TrimLeadingSpace = true
	r.FieldsPerRecord = -1
	r.ReuseRecord = true

	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	return records
}

func ArrayToCsv(array [][]string) string {
	f, err := os.Create("data/archivo.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // UTF-8 BOM

	w := csv.NewWriter(f)
	w.WriteAll(array)
	w.Flush()

	return "data/archivo.csv"
}
