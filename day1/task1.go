package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func getBoundsFromFile(filename string) (high, low int) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}

	low = 1000
	high = 1
	var scanner = bufio.NewScanner(bytes.NewReader(file))
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text()[1:])
		switch {
		case num < low:
			low = num
		case num > high:
			high = num
		default:
			continue
		}
	}
	return
}

var (
	input     = "2018day1input.txt"
	yieldchan = make(chan int, 1)
)

func main() {
	/*
		calibratedFreq, err := calibrate(input, yieldchan)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(calibratedFreq)
	*/
	// Task2
	highb, lowb := getBoundsFromFile(input)

	var v map[int]bool = make(map[int]bool)

	var device = NewWatchDevice()
}
