package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open(os.Args[1])
	scanner := bufio.NewScanner(file)
	var arr []int64
	lines := 0
	var N int
	for scanner.Scan() {
		text := scanner.Text()
		N = len(text) - 1
		v, _  := strconv.ParseInt(text, 2, 64)
		arr = append(arr, v)
		lines++
	}
	pos := N
	temp := make([]int64, lines)
	copy(temp, arr)
	for len(arr) != 1 {
		newArr1 := filter(arr, func(i int) bool {
			return 1<<pos&arr[i] == 1<<pos
		})
		newArr2 := filter(arr, func(i int) bool {
			return 1<<pos&arr[i] != 1<<pos
		})
		if len(newArr1) >= len(newArr2) {
			arr = newArr1
		} else {
			arr = newArr2
		}
		pos--
	}
	o2 := arr[0]
	arr = temp
	pos = N
	for len(arr) != 1 {
		newArr1 := filter(arr, func(i int) bool {
			return 1 << pos & arr[i] == 1 << pos
		})
		newArr2 := filter(arr, func(i int) bool {
			return 1 << pos & arr[i] != 1 << pos
		})
		if len(newArr1) < len(newArr2) {
			arr = newArr1
		} else {
			arr = newArr2
		}
		pos--
	}
	co2 := arr[0]
	fmt.Println(o2, co2, o2 * co2)
}

func filter(arr []int64, predicate func(int)bool) []int64 {
	var result []int64
	for i, v:= range arr {
		if predicate(i) {
			result = append(result, v)
		}
	}
	return result
}