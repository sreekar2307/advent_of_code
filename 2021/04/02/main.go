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
	scanner.Scan()
	text := scanner.Text()
	var draws []int
	{
		d := strings.Split(text,  ",")
		for _, v:= range d {
			i, _ := strconv.Atoi(v)
			draws = append(draws, i)
		}
	}
	k := 0
	var boards [][5][5]int
	var drawn [][5][5]bool
	var drawnMp []map[int][2]int
	for scanner.Scan() {
		boards = append(boards, [5][5]int{})
		drawn = append(drawn, [5][5]bool{})
		drawnMp = append(drawnMp, make(map[int][2]int))
		for i:=0;i<5;i++{
			scanner.Scan()
			text := scanner.Text()
			d := strings.Split(text,  " ")
			j := 0
			for _, v:= range d {
				if v == "" {
					continue
				}
				in, _ := strconv.Atoi(v)
				boards[k][i][j] = in
				drawnMp[k][in] = [2]int{i, j}
				j++
			}
		}
		k++
	}
	wonDraw := 0
	wonBoard := 0
	wonBoards := make(map[int]bool)
	label:
	for _, draw := range draws{
		for j := range boards {
			if v, ok := drawnMp[j][draw]; ok {
				drawn[j][v[0]][v[1]] = true
			}
		}
		for j :=  range boards {
			if _, ok := wonBoards[j]; ok {
				continue
			}
			if won(drawn[j]) {
				wonDraw = draw
				wonBoard = j
				wonBoards[j] = true
				if len(wonBoards) == len(boards) {
					break label
				}
			}
		}
	}
	sum := 0
	for i:=0;i<5;i++ {
		for j:=0;j<5;j++ {
			if !drawn[wonBoard][i][j]{
				sum += boards[wonBoard][i][j]
			}
		}
	}
	fmt.Println(sum*wonDraw)

}

func won(drawn [5][5]bool) bool {
	return wonRow(drawn, 0) || wonRow(drawn, 1) || wonRow(drawn, 2) || wonRow(drawn, 3) || wonRow(drawn, 4) ||
		wonCol(drawn, 0) || wonCol(drawn, 1) || wonCol(drawn, 2) || wonCol(drawn, 3) || wonCol(drawn, 4)
}

func wonRow(drawn [5][5]bool , row int) bool {
	return drawn[row][0] && drawn[row][1] && drawn[row][2] && drawn[row][3] && drawn[row][4]
}

func wonCol(drawn [5][5]bool , col int) bool {
	return drawn[0][col] && drawn[1][col] && drawn[2][col] && drawn[3][col] && drawn[4][col]
}