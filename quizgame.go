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
		log.Fatal("Dosyayı Bulamadım"+path, err) // in case of error it would return error
	}
	defer f.Close() //its close when func has done  because of "defer"

	csvReader := csv.NewReader(f) //reading
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Parse edemedim"+path, err)
	}
	problems := []Problem{}        //its our struct which holds all values
	for _, line := range records { //we ititrate over records to add problems array which is type is Problem
		problems = append(problems, Problem{line[0], line[1]})
	}
	return problems

}
func main() {
	records1 := readFromCsv("problems.csv") //pathway its static but I will change with argv
	score := 0
	for i, problem := range records1 {
		userAnswer := "" //to initilaze we have used decoy variable
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)
		fmt.Scan(&userAnswer) //Takes input
		if userAnswer == problem.answer {
			score++
		}
	}

	fmt.Printf("You scored %d out of %d.", score, len(records1)) //Score

}
