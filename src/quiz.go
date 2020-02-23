package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var input string
	var good, bad int

	file, err := os.Open("./quiz/problems.csv")
	check(err)
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		check(err)
		fmt.Println(record[0])
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println("Error: Empty line, please answer")
			bad++
			continue
		}
		if input == record[1] {
			good++
		} else {
			bad++
		}
	}
	fmt.Println("There were exactly", good+bad, "questions", "\nGood answer:", good, "\nBad answer:", bad)
}
