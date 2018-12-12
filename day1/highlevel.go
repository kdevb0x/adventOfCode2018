// This module contains packages build specifically to solve 2018's puzzles. By-and-for: kdevb0x
// more info about adventofcode can be found at https://adventofcode.com
//--- Copyright 2018 kdd | Licenced under MIT

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type FreqModule struct {
	lowest, highest int
	Visited         map[int]bool
}

func (p *FreqModule) scanForCollision(candidates chan int) (match <-chan int) {
	var matchChan = make(chan int, 1)
workque:
	for i := range candidates {
		if present, ok := p.Visited[i]; !ok {
			p.Visited[i] = false
			break workque
		} else {
			switch present {
			case false:
				p.Visited[i] = true
				break workque
			case true:
				matchChan <- i
				close(matchChan)
				break workque
			}

		}
	}
	close(candidates)
	return matchChan
}

func (p *FreqModule) calibrateFromReader(b io.Reader, rolloverFreq int) (current chan int, err error) {
	var currentFreq = make(chan int)
	var scanner = bufio.NewScanner(b)
	var freq int = 0

	if rolloverFreq == 0 {
	}
	freq = rolloverFreq

	for {
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
	}
	close(currentFreq)

	return currentFreq, nil

}

type WatchDevice struct {
	StarCount int
	Freqs     *FreqModule
}

func NewWatchDevice() *WatchDevice {
	var m map[int]bool = make(map[int]bool)
	fm := &FreqModule{
		highest: 0,
		lowest:  0,
		Visited: m,
	}
	watch := &WatchDevice{
		StarCount: 0,
		Freqs:     fm,
	}
	return watch
}

func (d *WatchDevice) CalibrateFreqs() error {
	caldata, err := ReadFileINTOBuffer(input)
	if err != nil {
		log.Println(err)
		return err
	}

	for {
		currentInt, _ := d.Freqs.calibrateFromReader(caldata, 0)
		colide := d.Freqs.scanForCollision(currentInt)
		select {
		case <-currentInt:
		case i := <-colide:
			fmt.Printf("the first freq to show up twice was %s\n", i)
		}
	}
	return nil

}

// ReadFile reads all data from 'filename' into a new bytes.Buffer and
// then closes the file. Any err other than nil returns nil buf
func ReadFileINTOBuffer(filename string) (*bytes.Buffer, error) {
	var b []byte
	buf := bytes.NewBuffer(b)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print("%s\n", err)
		return nil, err
	}
	if read, err := buf.Write(file); err != nil || len(file) != read {
		log.Println(err)
		return nil, err
	}
	return buf, nil
}
