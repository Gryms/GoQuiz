package main

import (
	"encoding/csv"
	"flag"
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

	path := flag.String("path", "./quiz/problems.csv", "Path to the quiz file you want to use")
	// timer := flag.Int("timer", 30, "Time in second that you have to answer all the questions from the quiz")

	flag.Parse()
	file, err := os.Open((*path))
	check(err)
	reader := csv.NewReader(file)
	for i := 0; i < 100; i++ {
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
