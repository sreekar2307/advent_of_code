package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("error while opening the input file %v %v\n", err, os.Args[1])
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var nums []int
	for scanner.Scan() {
		txt := scanner.Text()
		splits := strings.Split(txt, ",")
		nums = convertToIntArr(splits)
	}
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	for i := 0; i < 256; i++ {
		mpNew := make(map[int]int)
		for k, v := range mp {
			if k == 0 {
				mpNew[6] += v
				mpNew[8] += v
			} else {
				mpNew[k-1] += v
			}
		}
		mp = mpNew
	}
	var sum int64
	for _, v := range mp {
		sum += int64(v)
	}
	fmt.Println(sum)
}

func convertToIntArr(arr []string) []int {
	var result []int
	for _, v := range arr {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}
