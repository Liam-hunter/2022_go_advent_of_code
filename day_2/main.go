package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// first col: A=rock, B=paper, C=scissors
// secon col: X=rock, Y=paper, Z=scissors

// shape selection: rock=1, paper=2, scissors=3
// outcome: loss=0, 3=draw, 6=win

//rock > scissors
//rock < paper
//paper < scissors

func main() {
	input, err := readFile("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open input file: %v", err)
	}

	//input = "A Y\nB X\nC Z"

	total := 0
	for _, s := range strings.Split(input, "\n") {
		total += calculator(strings.Split(s, " "))
	}
	fmt.Println(total)

}

func readFile(f string) (string, error) {
	fp, err := os.Open(f)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	var b strings.Builder
	io.Copy(&b, fp)
	return b.String(), nil
}

func calculator(s []string) int {
	if len(s) == 2 {
		switch s[0] {
		case "A":
			//rock
			switch s[1] {
			case "X":
				return 4
			case "Y":
				return 8
			case "Z":
				return 3
			}
		case "B":
			//paper
			switch s[1] {
			case "X":
				return 1
			case "Y":
				return 5
			case "Z":
				return 9
			}
		case "C":
			//scissors
			switch s[1] {
			case "X":
				return 7
			case "Y":
				return 2
			case "Z":
				return 6
			}
		}
	}
	return 0
}
