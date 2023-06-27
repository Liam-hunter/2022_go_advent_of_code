package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	input, err := readFile("input.txt")
	if err != nil {
		log.Fatalf("Couldn't open input file: %v", err)
	}

	total := 0
	for _, s := range strings.Split(input, "\n") {
		c := strings.Split(s, " ")
		if len(c) == 2 {
			total += calculator(c[0], c[1])
		}
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

func calculator(a string, b string) int {
	if len(a) == 1 && len(b) == 1{
		switch a {
		case "A":
			//rock
			switch b {
			case "X":
				return 4
			case "Y":
				return 8
			case "Z":
				return 3
			}
		case "B":
			//paper
			switch b {
			case "X":
				return 1
			case "Y":
				return 5
			case "Z":
				return 9
			}
		case "C":
			//scissors
			switch b {
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
