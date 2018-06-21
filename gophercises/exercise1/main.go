package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type Quiz struct {
	Problem  string
	Solution string
}

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	data, err := openCSV("problems.csv")
	if err != nil {
		log.Fatal(err)
	}
	startQuiz(data)

}

func openCSV(filename string) ([][]string, error) {
	csvFile, _ := os.Open(filename)
	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	return lines, err

}

func startQuiz(data [][]string) {
	var numCorrect int
	var numQuestions int

	for _, line := range data {
		data := Quiz{
			Problem:  line[0],
			Solution: line[1],
		}
		fmt.Printf("%v=", data.Problem)

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()

		if text == data.Solution {
			fmt.Println("Correct!")
			numCorrect++
		} else {
			fmt.Println("Incorrect!")
		}
		numQuestions++
	}
	fmt.Printf("There were %v questions asked, and you got %v questions correct!\n", numQuestions, numCorrect)
}
