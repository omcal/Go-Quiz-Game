package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Problem struct {
	question string
	answer   string
}

func readFromCsv(path string) []Problem {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Dosyayı Bulamadım"+path, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Parse edemedim"+path, err)
	}
	problems := []Problem{}
	for _, line := range records {
		problems = append(problems, Problem{line[0], line[1]})
	}
	return problems

}
func main() {
	records1 := readFromCsv("problems.csv")
	score := 0
	for i, problem := range records1 {
		userAnswer := ""
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		fmt.Scan(&userAnswer)
		if userAnswer == problem.answer {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d.", score, len(records1))

}
