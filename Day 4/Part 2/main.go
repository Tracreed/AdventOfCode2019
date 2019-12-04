package main

import (
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
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
	numbers := strings.Split(string(byteValue), "-")
	in1, err := strconv.Atoi(numbers[0])
	check(err)
	in2, err := strconv.Atoi(numbers[1])
	check(err)
	total := 0
	for i := in1; i <= in2; i++ {
		st := strconv.Itoa(i)
		sts := strings.Split(st, "")
		ints := make([]int, len(sts))
		for l, k := range sts {
			ints[l], _ = strconv.Atoi(string(k))
		}
		if sort.IntsAreSorted(ints) && IntGotDoubles(ints) {
			total++
		}
	}
	fmt.Println(total)
}

func IntGotDoubles(x []int) bool {
	m := make(map[int]int)
	for _, v := range x {
		m[v]++
	}
	for _, n := range m {
		if n > 1 && n < 3 {
			return true
		}
	}
	return false
}