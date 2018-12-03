package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	foundchan = make(chan int, 1)
	lastpool  []int
)

func searchLast(match int, last []int, foundchan chan int) {
	for _, val := range last {
		if match == val {
			foundchan <- val
		}
	}
	foundchan <- 0
}

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
			if strings.Contains(strconv.Itoa(lastpool[:]), string(freq)) {
				duperr := fmt.Errorf("duplicate found before before parsing finished")
				return freq, duperr
			}
			lastpool = append(lastpool, freq)
		}
		if strings.ContainsRune(scanner.Text(), '+') {
			change, err := strconv.Atoi(scanner.Text()[1:])
			if err != nil {
				log.Println(err)
				continue
			}
			freq += change
			lastpool = append(lastpool, freq)
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
