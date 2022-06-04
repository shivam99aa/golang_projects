package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func askQuestion(trackAnswer chan bool, record []string) {
	var answer string
	fmt.Println(record[0])
	fmt.Scanln(&answer)

	if answer == record[1] {
		trackAnswer <- true
	} else {
		trackAnswer <- false
	}
}

func trackTimeOut(trackTime chan bool, timeOutValue int) {
	time.Sleep(time.Duration(timeOutValue) * time.Second)
	trackTime <- true
}

func main() {
	csvFileName := flag.String("filename", "problems.csv", "filename containing problems")
	timeOutValue := flag.Int("timeout", 30, "timeout for answering question")
	flag.Parse()

	var correct, totalQuestions int
	trackAnswer := make(chan bool)
	trackTime := make(chan bool)

	f, err := os.Open(*csvFileName)

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	fmt.Println("Timeout value is ", *timeOutValue, ". Press Enter to start the quiz.")
	fmt.Scanln()
	go trackTimeOut(trackTime, *timeOutValue)
	for {
		record, err := r.Read()

		if err == io.EOF {
			fmt.Println("Total Questions =>", totalQuestions)
			fmt.Println("Correct Questions =>", correct)
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		totalQuestions++
		go askQuestion(trackAnswer, record)

		select {
		case q := <-trackAnswer:
			if q == true {
				correct++
			}
		case <-trackTime:
			fmt.Println("Total Questions =>", totalQuestions)
			fmt.Println("Correct Questions =>", correct)
			return
		}

	}
}
