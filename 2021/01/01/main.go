package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	arr := make([]int, 0)
	for scanner.Scan() {
		integer, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, integer)
	}
	increaseCount := 0
	for i:=1;i<len(arr);i++ {
		if arr[i-1] < arr[i] {
			increaseCount++
		}
	}
	fmt.Println(increaseCount)
}
