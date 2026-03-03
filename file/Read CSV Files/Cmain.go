package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// CSV ফাইল ওপেন
	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// CSV রিডার তৈরি
	reader := csv.NewReader(file)

	// সব Rows পড়া (slice of records)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// রেসাল্ট লুপ করে দেখা
	for _, row := range records {
		fmt.Println(row)
	}
}
