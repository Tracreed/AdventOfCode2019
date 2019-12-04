package main

import (
	"os"
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
	"math"
	"sort"
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
	cross := []float64{}
	for i, w := range wires {
		s := strings.Split(w, ",")
		fmt.Println("Wire " + strconv.Itoa(i + 1) + ":")
		x, y  := 11500, 11500
		for _, v := range s {
			direction := v[0:1]
			steps, _ := strconv.Atoi(v[1:])

			fmt.Printf("Direction: %s Steps: %d\n", direction, steps)
			switch(direction) {
			case "U":
				for l := 0; l < steps; l++ {
					if pan[x - 1][y] != 0 && pan[x - 1][y] != i + 1 {
						fmt.Println("lines cross")
						d := math.Abs(float64(x - 1 - 11500)) + math.Abs(float64(y - 11500))
						cross = append(cross, d)
					}
					pan[x - 1][y] = i + 1
					x -= 1
					
				}
			case "D":
				for l := 0; l < steps; l++ {
					if pan[x + 1][y] != 0 && pan[x + 1][y] != i + 1 {
						fmt.Println("lines cross")
						d := math.Abs(float64(x + 1 - 11500)) + math.Abs(float64(y - 11500))
						cross = append(cross, d)
					}
					pan[x + 1][y] = i + 1
					x += 1
				}
			case "R":
				for l := 0; l < steps; l++ {
					if pan[x][y + 1] != 0 && pan[x][y + 1] != i + 1 {
						fmt.Println("lines cross")
						d := math.Abs(float64(x - 11500)) + math.Abs(float64(y + 1 - 11500))
						cross = append(cross, d)
					}
					pan[x][y + 1] = i + 1
					y += 1
				}
			case "L":
				for l := 0; l < steps; l++ {
					if pan[x][y - 1] != 0 && pan[x][y - 1] != i + 1 {
						fmt.Println("lines cross")
						d := math.Abs(float64(x - 11500)) + math.Abs(float64(y - 1 - 11500))
						cross = append(cross, d)
					}
					pan[x][y - 1] = i + 1
					y -= 1
				}
			}
			fmt.Println(x, y)
		}
	}
	sort.Float64s(cross)
	fmt.Println(cross)
}