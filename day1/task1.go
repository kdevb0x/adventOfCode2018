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

type intpool struct {
	lowest, highest int
	visited         map[int]bool
}

func (p *intpool) scanForCollision(candidates chan int) (match int, found bool) {
workque:
	for i, ok := range <-candidates; ok ok != false {
		present, seen := p.visited[i]
		if seen == false {
			p.visited[i] = false
			break workque
		} else {
			switch present {
			case false:
				p.visited[i] = true
				break workque
			case true:
				return i, true
			}

		}
	}
	return 0, false
}

func getBounds(filename string) (high int, low int) {
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

func calibrate(inputFilename string, currentFreq chan int, rolloverFreq int) (int, error) {
	file, err := os.Open(inputFilename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	var scanner = bufio.NewScanner(file)
	var freq int

	if rolloverFreq != 0 {
		freq = rolloverFreq
	}
	freq = 0

	for scanner.Scan() {
		if strings.ContainsRune(scanner.Text(), '-') {
			change, err := strconv.Atoi(scanner.Text()[1:])
			if err != nil {
				log.Println(err)
				continue
			}
			freq -= change
			/*
					if strings.Contains(strconv.Itoa(lastpool[:]), string(freq)) {
						duperr := fmt.Errorf("duplicate found before before parsing finished")
						return freq, duperr
					}
					lastpool = append(lastpool, freq)
				}
			*/
		}
		if strings.ContainsRune(scanner.Text(), '+') {
			change, err := strconv.Atoi(scanner.Text()[1:])
			if err != nil {
				log.Println(err)
				continue
			}
			freq += change
		}
		currentFreq <- freq
	}
	return freq, nil

}

func main() {
	/*
		calibratedFreq, err := calibrate(input, yieldchan)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(calibratedFreq)
	*/
	// Task2
	highb, lowb := getBounds(input)
	var v map[int]bool = make(map[int]bool)
	freqpool := &intpool{
		highest: highb,
		lowest:  lowb,
		visited: v,
	}

	// run the calibration continuously until the a collision is found
	hit := func() int {
		var hit int
		for n, found := freqpool.scanForCollision(yieldchan); found == false; {
			for final, _ := calibrate(input, yieldchan, 0); final >= 0; {
				final, _ = calibrate(input, yieldchan, final)
				hit = n
			}

		}
		return hit
	}
	var d = hit()
	fmt.Printf("The first freq to be seen twice was: %d\n", d)

}
