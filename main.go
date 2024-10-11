package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// Открываем CSV файл
	file, err := os.Open("sample.csv")
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

	// Выводим данные
	for _, record := range records {
		fmt.Println(record)
	}
}
