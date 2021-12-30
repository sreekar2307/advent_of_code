package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("error while opening the input file %v %v\n", err, os.Args[1])
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var counts []int
	lines := 0
	for scanner.Scan() {
		text := scanner.Text()
		for i, v := range text {
			if len(counts) < i+1 {
				counts = append(counts, 0)
			}
			if v == '1' {
				counts[i]++
			}
		}
		lines++
	}
	alpha := ""
	gamma := ""
	for _, v := range counts {
		if lines%2 == 0 && v >= lines/2 {
			alpha += "1"
			gamma += "0"
		} else if lines%2 != 0 && v >= lines/2+1 {
			alpha += "1"
			gamma += "0"
		} else {
			alpha += "0"
			gamma += "1"
		}
	}
	a, _ := strconv.ParseInt(alpha, 2, 64)
	b, _ := strconv.ParseInt(gamma, 2, 64)
	fmt.Println(a, b, a*b)
}
