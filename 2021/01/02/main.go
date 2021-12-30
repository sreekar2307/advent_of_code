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
	sumPrev :=arr[0]+arr[1]+arr[2]
	for i:=1;i<len(arr)-2;i++ {
		currSum := arr[i]+arr[i+1]+arr[i+2]
		if sumPrev < currSum {
			increaseCount++
		}
		sumPrev = currSum
	}
	fmt.Println(increaseCount)
}
