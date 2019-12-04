package main

import (
	"os"
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
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

	temp := strings.Split(string(byteValue), "\n")
	fmt.Println(temp)
	total := 0
	total2 := 0
	for i := 0; i < len(temp); i++ {
		if len(temp[i]) == 0 {
			continue
		}
		numb, err := strconv.Atoi(temp[i])
		check(err)
		total += numb / 3 - 2
		total2 += numb / 3 - 2
		temptotal := numb / 3 - 2
		for temptotal / 3 - 1 > 0 {
			total2 += temptotal / 3 - 2
			temptotal = temptotal / 3 - 2
		}
	}
	fmt.Println(total)
	fmt.Println(total2)
}