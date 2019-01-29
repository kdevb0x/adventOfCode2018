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
	var num int

	rbuff := bufio.NewScanner(file)
	for rbuff.Scan() {
		n, _ := strconv.Atoi(rbuff.Text())
		num += n
	}
	if err := rbuff.Err(); err != nil {
		log.Fatal(err)
	}
	return num
}

func main() {
	numb := calibrate()
	fmt.Println(numb)
}
