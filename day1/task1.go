package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calibrate(inputFilename string) (int, error) {
	file, err := os.Open(inputFilename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)
	var freq = 0

	for scanner.Scan() {
		if strings.ContainsRune(scanner.Text(), '-') {
			change, err := strconv.Atoi(scanner.Text()[1:])
			if err != nil {
				log.Println(err)
				continue
			}
			freq -= change
		}
		if strings.ContainsRune(scanner.Text(), '+') {
			change, err := strconv.Atoi(scanner.Text()[1:])
			if err != nil {
				log.Println(err)
				continue
			}
			freq += change
		}
	}
	return freq, nil

}

func main() {
	calibratedFreq, err := calibrate("2018day1input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(calibratedFreq)
}
