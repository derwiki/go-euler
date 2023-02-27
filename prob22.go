package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
)

func alphabeticalValue(xs string) int {
	acc := 0
	for i := 0; i < len(xs); i++ {
		acc += int(xs[i] - 'A' + 1)
	}
	return acc
}

func main() {
	file, err := os.Open("p022_names.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Set the comma separator
	reader.Comma = ','

	// Read all the records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create an array to hold the file names
	var names []string

	// Loop through each record
	for _, record := range records {
		// Loop through each field in the record
		for _, field := range record {
			// Remove the quotes around the file name
			name := strings.Trim(field, "\"")
			names = append(names, name)
		}
	}

	// Sort the file names
	sort.Strings(names)

	acc := 0
	for i, name := range names {
		// fmt.Println(i, name, alphabeticalValue(name))
		acc += alphabeticalValue(name) * (i + 1)
	}
	fmt.Println("acc", acc)
}
