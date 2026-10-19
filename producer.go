package main

import (
	"encoding/csv"
	"os"
)

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func loadRecipient(filePath string, ch chan Recipient) error {

	defer close(ch)

	f, err := os.Open(filePath)
	checkNilError(err)
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	checkNilError(err)

	for _, record := range records[1:] {
		// fmt.Printf("User name is %v with email %v\n", record[0], record[1])
		// fmt.Println(record)

		ch <- Recipient{
			Name:  record[0],
			Email: record[1],
		}
	}

	return nil
}
