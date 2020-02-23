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
	file, err := os.Open("./quiz/test.txt")
	check(err)
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		check(err)
		fmt.Println(record[0])
	}
}
