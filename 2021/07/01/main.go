package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	sort.Ints(nums)
	minAns := math.MaxInt32
	max := nums[len(nums)-1]
	for i := 0; i <= max; i++ {
		sum := 0
		for _, v := range nums {
			n := abs(i - v)
			sum += n
		}
		minAns = min(sum, minAns)
	}
	fmt.Println(minAns)
}

func convertToIntArr(arr []string) []int {
	var result []int
	for _, v := range arr {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
