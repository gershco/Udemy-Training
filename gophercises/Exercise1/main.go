package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Set up flag(s)
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question/answer'")

	flag.Parse()

	// Opens the named file
	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("CSV file %s failed to open.\n", *csvFilename))
	}

	r := csv.NewReader(file)

	// Parse the whole file into a 2D slice.
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the CSV file")
	}

	problems := parseLines(lines)

	score := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			score++
		}
		fmt.Printf("You scored %d out of %d.\n", score, len(problems))
	}

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
