package main

import (
	"encoding/csv"
	"fmt"
	"github.com/sqweek/dialog"
	"log"
	"os"
)

func main() {
	// Открываем диалоговое окно для выбора файла
	filename, err := dialog.File().Filter("CSV Files", "csv").Title("Выберите CSV файл").Load()
	if err != nil {
		log.Fatal(err)
	}

	// Открываем выбранный файл
	file, err := os.Open(filename)
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

	// Выводим содержимое CSV
	for _, record := range records {
		fmt.Println(record)
	}
}
