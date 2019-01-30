package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func calibrate() int {
	file, err := os.Open("2018day1input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := bufio.NewScanner(file)

	var (
		num int // current total
		// nums []int  // running list of totals
		numchan   = make(chan int, 1000)
		foundchan = make(chan int, 1)
	)

	go checkDup(numchan, foundchan) // range through nums and check

	for r.Scan() {
		n, err := strconv.Atoi(r.Text())
		if err != nil {
			log.Println(err)
		}
		num += n
		numchan <- num
		if d := <-foundchan; d > 0 {
			close(numchan)
			break
		}
	}
	return <-foundchan
}

func checkDup(input chan int, found chan int) {
	dups := map[int]bool{0: true}
	for n := range input {
		if dups[n] {
			found <- n
			return
		}
		dups[n] = true
	}

}

func main() {
	numb := calibrate()
	fmt.Println(numb)
}
