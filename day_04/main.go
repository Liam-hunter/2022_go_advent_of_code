package main

import (
	"fmt"
	"log"
	"os"
)

const (
	lineFeed = 10
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	length := 0
	count := 0
	for index, byte := range bytes {
		if byte == lineFeed {
			item := itemFinder(bytes[index-length:index], length)
			if item > 96 && item < 123 {
				count += int(item) - 96
			}
			if item > 64 && item < 91 {
				count += int(item) - 38
			}
			length = 0
			continue
		}
		length++
	}

	fmt.Println(count)

}

func itemFinder(b []byte, l int) byte {
	if (l % 2) != 0 {
		return byte(0)
	}

	for _, first := range b[0:(l / 2)] {
		for _, second := range b[l/2:] {
			if first == second {
				return first
			}
		}
	}
	return byte(0)
}
