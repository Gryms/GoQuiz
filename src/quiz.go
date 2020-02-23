package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func end(good, bad int) {
	fmt.Println("There were exactly", good+bad, "questions", "\nGood answer:", good, "\nBad answer:", bad)
}

func main() {
	var input string
	var good, bad int

	path := flag.String("path", "./quiz/problems.csv", "Path to the quiz file you want to use")
	timer := flag.Int("timer", 30, "Time in second that you have to answer all the questions from the quiz")

	flag.Parse()
	file, err := os.Open((*path))
	check(err)
	reader := csv.NewReader(file)

	watch := time.NewTimer(time.Duration(*timer) * time.Second)
	go func() {
		<-watch.C
		fmt.Println("Time's up !")
		end(good, bad)
		os.Exit(0)
	}()

	for i := 0; i < 100; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		check(err)
		fmt.Println("What's the answer to:", record[0], "?")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println("Error: Empty line, please answer")
			bad++
			continue
		}
		if input == record[1] {
			fmt.Println("That's right !")
			good++
		} else {
			fmt.Println("Huh... actually, no.")
			bad++
		}
	}
	end(good, bad)
}
