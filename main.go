package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	q string
	a string
}

//Part1 of the Quiz game
func main() {

	//Reading the csv filename using flags
	csvFileName := flag.String("csv", "problems.csv", "A csv fle containing the quiz in <question,answer> format")
	flag.Parse()

	//Opening the file
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Could not open csv file: %s\n", *csvFileName))
	}

	//Reading the whole file at a time as csv file wont be large
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Could not read the csv file" + err.Error()))
	}

	//Getting the problems
	problems := getProblems(lines)

	//Printing each problem
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.a {
			correct++
		}
	}
	//Final Result
	fmt.Printf("You scored %d out of %d!\n", correct, len(problems))

}

//getProblems parses the the file data into a list of problem struct
func getProblems(lines [][]string) []problem {
	//Its better to make a defnes slice and not use append as we already know the length. More performance
	result := make([]problem, len(lines))
	for i, line := range lines {
		result[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), //trimming extra spaces from csv
		}
	}
	return result
}

//exit prints a message and exits the program
func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
