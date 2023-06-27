package main

import (
	"os"
	"fmt"
	//"strings"
	//"io"
	"log"
)

const ( 
	rock = 65 //rock A
	paper = 66 //paper B
	scissors = 67 //scissors C

	lose = 88 //lose x
	draw = 89 //draw y
	win = 90 //win z
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("couldn't load file: %v", err)
	}

	total := 0
	for i, byte := range bytes {
		switch byte{
		case lose:
			switch bytes[i-2] {
			case rock:
				total += 3
			case paper:
				total += 1
			case scissors:
				total += 2
			}
		case draw:
			switch bytes[i-2] {
			case rock:
				total += 4
			case paper:
				total += 5
			case scissors:
				total += 6
			}
		case win:
			switch bytes[i-2] {
			case rock:
				total += 8
			case paper:
				total += 9
			case scissors:
				total += 7
			}
		}
	}
	fmt.Println(total)
}

