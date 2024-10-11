package main

import (
	"encoding/csv"
	"fmt"
	"log"
)

func main() {
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Читаем содержимое CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		fmt.Println(record)
	}
}

// Test commit
