package main

import (
	"fmt"
	"os"
)

const (
	lineFeed = 10
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(1)
	}

	lineBytes := make([][]byte, 3)
	lines := 0
	lastLF := 0
	count := 0

	for index, byte := range bytes {
		switch byte {
		case lineFeed:
			lineBytes[lines] = bytes[lastLF+1 : index]
			if lines == 2 {
				//something
				item := itemFinder(lineBytes)
				if item > 96 && item < 123 {
					count += int(item) - 96
				}
				if item > 64 && item < 91 {
					count += int(item) - 38
				}
				lines = 0
			} else {
				lines++
			}
			lastLF = index
		}
	}
	fmt.Println(count)

}

func itemFinder(b [][]byte) byte {
	for _, byte := range b[0] {
		if finder(b[1], byte) && finder(b[2], byte) {
			return byte
		}
	}
	return byte(0)
}

func finder(bytes []byte, byteToFind byte) bool {
	for _, byte := range bytes {
		if byte == byteToFind {
			return true
		}
	}
	return false
}
