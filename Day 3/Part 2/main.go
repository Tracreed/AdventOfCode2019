package main

import (
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"math"
	"sort"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	inputFile, err := os.Open("./input")
	check(err)

	byteValue, _ := ioutil.ReadAll(inputFile)
	inputFile.Close()

	wires := strings.Split(string(byteValue), "\n")
	pan := [23000][23000]int{}
	cross := make(map[float64][]int)
	path := [][][]int{[][]int{},[][]int{}}
	for i, w := range wires {
		s := strings.Split(w, ",")
		x, y  := 11500, 11500
		for _, v := range s {
			direction := v[0:1]
			steps, _ := strconv.Atoi(v[1:])
			switch(direction) {
			case "U":
				for l := 0; l < steps; l++ {
					if pan[x - 1][y] != 0 && pan[x - 1][y] != i + 1 {
						d := math.Abs(float64(x - 1 - 11500)) + math.Abs(float64(y - 11500))
						cross[d] = []int{x - 1, y}
					}
					pan[x - 1][y] = i + 1
					x -= 1
					p := []int{x, y}
					path[i] = append(path[i], p)
				}
			case "D":
				for l := 0; l < steps; l++ {
					if pan[x + 1][y] != 0 && pan[x + 1][y] != i + 1 {
						d := math.Abs(float64(x + 1 - 11500)) + math.Abs(float64(y - 11500))
						cross[d] = []int{x + 1, y}
					}
					pan[x + 1][y] = i + 1
					x += 1
					p := []int{x, y}
					path[i] = append(path[i], p)
				}
			case "R":
				for l := 0; l < steps; l++ {
					if pan[x][y + 1] != 0 && pan[x][y + 1] != i + 1 {
						d := math.Abs(float64(x - 11500)) + math.Abs(float64(y + 1 - 11500))
						cross[d] = []int{x, y + 1}
					}
					pan[x][y + 1] = i + 1
					y += 1
					p := []int{x, y}
					path[i] = append(path[i], p)
				}
			case "L":
				for l := 0; l < steps; l++ {
					if pan[x][y - 1] != 0 && pan[x][y - 1] != i + 1 {
						d := math.Abs(float64(x - 11500)) + math.Abs(float64(y - 1 - 11500))
						cross[d] = []int{x , y - 1}
					}
					pan[x][y - 1] = i + 1
					y -= 1
					p := []int{x, y}
					path[i] = append(path[i], p)
				}
			}
		}
	}
	pl := []int{}
	for _, g := range cross {
		pl = append(pl, walk(path[0], g[0], g[1]) + walk(path[1], g[0], g[1]) + 2)
	}
	sort.Ints(pl)
	fmt.Println(pl[0])
}

func walk(path [][]int, gx int, gy int) int {
	total := 0
	for _, v := range path {
		if v[0] == gx && v[1] == gy {
			return total
		}
		total++
	}
	return total
}