/* This module contains packages build specifically to solve 2018's puzzles
*  more info about adventofcode can be found at https://adventofcode.com
*
*  Copyright 2018 kdevb0x Ltd. All rights reserved
*  Use of this source code is governed by the BSD 3-Clause license
*  The full license text can be found in the LICENSE file.
 */

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type FreqModule struct {
	lowest, highest int

	// freq will set to false when seen the first time, true the 2nd.
	Visited map[int]bool
}

func (p *FreqModule) scanForCollision(candidates chan int) (match <-chan int) {
	var matchChan = make(chan int, 1)
workque:
	for i := range candidates {
		if present, ok := p.Visited[i]; !ok {
			p.Visited[i] = true
			break workque
		} else {
			switch present {
			case false:
				p.Visited[i] = true
				matchChan <- i
				match = matchChan
			}

		}
	}
	close(candidates)
	return matchChan
}

func (p *FreqModule) calibrateFromReader(b io.Reader, rolloverFreq int) (current chan int, err error) {
	currentFreq := make(chan int)
	current = currentFreq
	var scanner = bufio.NewScanner(b)
	var freq int

	if rolloverFreq == 0 {
	}
	freq = rolloverFreq

	for scanner.Scan() {

		// Subtract from the net freq
		if strings.ContainsRune(scanner.Text(), '-') {
			change, err := strconv.Atoi(scanner.Text()[1:])
			if err != nil {
				log.Println(err)
				continue
			}
			freq -= change
		}

		// Add to the net freq
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

	close(currentFreq)

	err = nil
	return

}

type WatchDevice struct {
	StarCount int
	Freqs     *FreqModule
}

// NewWatchDevice creates our time-traveling device 
func NewWatchDevice() *WatchDevice {
	var m = make(map[int]bool)
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

// CalibrateFreqs does what it's name implies, reading input data from a file.
func (d *WatchDevice) CalibrateFreqs() error {
	caldata, err := ReadFileINTOBuffer(input)
	if err != nil {
		log.Println(err)
		return err
	}

	currentInt, _ := d.Freqs.calibrateFromReader(caldata, 0)
	colide := d.Freqs.scanForCollision(currentInt)
	select {
	case <-currentInt:
	case i := <-colide:
		fmt.Printf("the first freq to show up twice was %d\n", i)
	}

	return nil

}

// ReadFileINTOBuffer reads all data from 'filename' into a new bytes.Buffer,
// (cleanly opening and closing the file) returning a pointer to the Buffer.
//
// NOTE: If any non-nil errors are encountered, Buffer will be a nil pointer!!
func ReadFileINTOBuffer(filename string) (*bytes.Buffer, error) {
	var b []byte
	buf := bytes.NewBuffer(b)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	if read, err := buf.Write(file); err != nil || len(file) != read {
		log.Println(err)
		return nil, err
	}
	return buf, nil
}
