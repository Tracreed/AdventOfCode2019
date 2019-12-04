package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

	temps := strings.Split(string(byteValue), ",")
	var temp = []int{}
	for _, i := range temps {
		j, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		temp = append(temp, j)

	}
	var newtemp = make([]int, len(temp))
	copy(newtemp, temp)
	for p1 := 1; p1 < 100; p1++ {
		for p2 := 1; p2 < 100; p2++ {
			newtemp[1] = p1
			newtemp[2] = p2
			res := runInt(newtemp)
			if res[0] == 19690720 {
				fmt.Println(100*p1 + p2)
			}
			copy(newtemp, temp)
		}
	}
}

func runInt(pr []int) []int {
	for i := 0; i < len(pr); i++ {
		k := i
		v := pr[k]
		switch v {
		case 1:
			v1 := pr[pr[k+1]]
			v2 := pr[pr[k+2]]
			pr[pr[k+3]] = v1 + v2
			i += 3
		case 2:
			v1 := pr[pr[k+1]]
			v2 := pr[pr[k+2]]
			pr[pr[k+3]] = v1 * v2
			i += 3
		case 99:
			return pr
		}
	}
	return pr
}
