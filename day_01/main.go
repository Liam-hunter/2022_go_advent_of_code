package main

import (
	"os"
	"io"
	"fmt"
	"log"
	"strconv"
	"strings"
	"main/calories"
)

func main() {
	s, err := parseFile("input.txt")
	if err != nil {
		log.Fatal(err)

	}

	maxCals := calories.NewMax()

	n := 0
	for _, s := range(s) {
		if s == "" {
			maxCals.SetMax(n)
			n = 0
			continue
		}
		if x, err := strconv.Atoi(s); err != nil {
			log.Fatal(err)
		} else {
			n += x
		}
	}

	fmt.Print("The maximum calories is: ")

	fmt.Println(maxCals.GetMax())

}


func parseFile(path string) ([]string, error) {
	fp, err := os.Open(path)
	defer fp.Close()
	if err != nil {
		return nil, err
	}

	var bytes strings.Builder

	_, err = io.Copy(&bytes, fp)
	if err != nil {
		return nil, err
	}

	return strings.Split(bytes.String(), "\n"), nil
}
