package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	x,y := 0,0
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " ")
		if splits[0] == "forward" {
			i, _ :=  strconv.Atoi(splits[1])
			x+=i
		}else if splits[0] == "up" {
			i, _ :=  strconv.Atoi(splits[1])
			y+=i
		}else {
			i, _ :=  strconv.Atoi(splits[1])
			y-=i
		}
	}
	fmt.Println(abs(x*y))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}