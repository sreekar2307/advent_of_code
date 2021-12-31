package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CoOrdinate struct {
	x, y, z int
}

func (co CoOrdinate) isNull() bool {
	return co.x == 0 && co.y == 0 && co.z == 0
}
func (co CoOrdinate) dot(co1 CoOrdinate) int {
	return co.x*co1.x + co.y*co1.y + co.z*co1.z
}
func (co CoOrdinate) cross(co1 CoOrdinate) CoOrdinate {
	return CoOrdinate{
		x: co.y*co1.z - co.z*co1.y,
		y: co.z*co1.x - co.x*co1.z,
		z: co.x*co1.y - co.y*co1.x,
	}
}

func (co CoOrdinate) String() string {
	return fmt.Sprintf("%d,%d", co.x, co.y)
}

type LineSegment struct {
	pointA CoOrdinate
	pointB CoOrdinate
}

func (ls LineSegment) minXCoOrdinate() int {
	return min(ls.pointA.x, ls.pointB.x)
}
func (ls LineSegment) maxXCoOrdinate() int {
	return max(ls.pointA.x, ls.pointB.x)
}

func (ls LineSegment) minYCoOrdinate() int {
	return min(ls.pointA.y, ls.pointB.y)
}
func (ls LineSegment) maxYCoOrdinate() int {
	return max(ls.pointA.y, ls.pointB.y)
}

func (ls LineSegment) vector() CoOrdinate {
	return CoOrdinate{
		x: ls.pointB.x - ls.pointA.x,
		y: ls.pointB.y - ls.pointA.y,
		z: ls.pointB.z - ls.pointA.z,
	}
}

func (ls LineSegment) covers(point CoOrdinate) bool {
	ab := ls.vector()
	ac := LineSegment{
		pointA: ls.pointA,
		pointB: point,
	}.vector()
	if !ab.cross(ac).isNull() {
		return false
	}
	kac := ab.dot(ac)
	kab := ab.dot(ab)
	return 0 <= kac && kac <= kab
}
func (ls LineSegment) parallelToXOrY() bool {
	return ls.pointB.x == ls.pointA.x || ls.pointA.y == ls.pointB.y
}
func (ls LineSegment) String() string {
	return fmt.Sprintf("%s -> %s", ls.pointA, ls.pointB)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("error while opening the input file %v %v\n", err, os.Args[1])
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []LineSegment
	for scanner.Scan() {
		txt := scanner.Text()
		splits := strings.Split(txt, ",")
		nums := convertToIntArr(splits)
		pa := CoOrdinate{
			x: nums[0],
			y: nums[1],
		}
		pb := CoOrdinate{
			x: nums[2],
			y: nums[3],
		}
		ls := LineSegment{
			pointA: pa,
			pointB: pb,
		}
		lines = append(lines, ls)
	}
	minX, maxX, minY, maxY := math.MaxInt32, math.MinInt32, math.MaxInt32, math.MinInt32
	for _, line := range lines {
		minX = min(minX, line.minXCoOrdinate())
		minY = min(minY, line.minYCoOrdinate())
		maxX = max(maxX, line.maxXCoOrdinate())
		maxY = max(maxY, line.maxYCoOrdinate())
	}
	greaterThan2 := 0
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			covers := 0
			for _, line := range lines {
				if line.parallelToXOrY() && line.covers(CoOrdinate{i, j, 0}) {
					covers++
				}
			}
			if covers >= 2 {
				greaterThan2++
			}
		}
	}
	fmt.Println(greaterThan2)
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
