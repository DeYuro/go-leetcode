package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	city := [][]int{}
	var coord string
	for n > 0 {
		fmt.Scanf("%s", &coord)
		spl := strings.Split(coord," ")
		x, _ := strconv.Atoi(spl[0])
		y, _ := strconv.Atoi(spl[1])
		city = append(city, []int{x,y})
	}

	var distance int
	fmt.Scanf("%d", &distance)

	var fromWhere string
	fmt.Scanf("%s", &fromWhere)
	spl := strings.Split(fromWhere," ")

	idxFrom, _ := strconv.Atoi(spl[0])
	idxWhere, _ := strconv.Atoi(spl[1])

	if idxFrom == idxWhere {
		fmt.Println(0)
		return
	}

	queue := [][2]int{}

	queue = append(queue,[2]int{idxFrom,0})

	visited := make(map[int]bool)
	visited[idxFrom] = true
	for len(queue) > 0 {
		arr := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if arr[0] == idxWhere {

			fmt.Println(arr[1])
			return
		}

		for i := 0; i < n; i++ {
			if maxDistance(city[i], city[arr[0]]) <= distance {
				if _, ok := visited[i]; !ok {
					queue = append(queue, [2]int{i,arr[1]})
					visited[i] = true
				}
			}
		}
	}

	fmt.Println(-1)
}


func maxDistance(from, where []int) int  {
	return abs(from[0] - where[0]) + abs(from[1] - from[2])
}

func abs(a int) int {
	if a > 0 {
		return a
	}

	return a * -1
}
