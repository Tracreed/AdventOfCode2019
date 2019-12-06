package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"bufio"
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
	runInt(temp)
}

func runInt(pr []int) []int {
	for i := 0; i < len(pr); i++ {
		k := i
		v := pr[k]
		//fmt.Println(pr)
		switch OpCode(v)["DE"] {
		case 1:
			v1 := 0
			v2 := 0
			if OpCode(v)["C"] == 1 {
				v1 = pr[k+1]
			} else {
				v1 = pr[pr[k+1]]
			}
			if OpCode(v)["B"] == 1 {
				v2 = pr[k+2]
			} else {
				v2 = pr[pr[k+2]]
			}
			pr[pr[k+3]] = v1 + v2
			i += 3
		case 2:
			v1 := 0
			v2 := 0
			if OpCode(v)["C"] == 1 {
				v1 = pr[k+1]
			} else {
				v1 = pr[pr[k+1]]
			}
			if OpCode(v)["B"] == 1 {
				v2 = pr[k+2]
			} else {
				v2 = pr[pr[k+2]]
			}
			pr[pr[k+3]] = v1 * v2
			i += 3
		case 3:
			in := readInput()
			v1 := pr[k+1]
			pr[v1] = in
			i += 1
		case 4:
			v1 := 0
			if OpCode(v)["C"] == 1 {
				v1 = pr[k+1]
			} else {
				v1 = pr[pr[k+1]]
			}
			fmt.Println(v1)
			i += 1
		case 5:
			v1 := 0
			v2 := 0
			if OpCode(v)["C"] == 1 {
				v1 = pr[k+1]
			} else {
				v1 = pr[pr[k+1]]
			}
			if OpCode(v)["B"] == 1 {
				v2 = pr[k+2]
			} else {
				v2 = pr[pr[k+2]]
			}
			if v1 != 0 {
				i = v2 - 1
			} else {
				i += 2
			}
		case 6:
			v1 := 0
			v2 := 0
			if OpCode(v)["C"] == 1 {
				v1 = pr[k+1]
			} else {
				v1 = pr[pr[k+1]]
			}
			if OpCode(v)["B"] == 1 {
				v2 = pr[k+2]
			} else {
				v2 = pr[pr[k+2]]
			}
			if v1 == 0 {
				i = v2 - 1
			} else {
				i += 2
			}
		case 7:
			v1 := 0
			v2 := 0
			if OpCode(v)["C"] == 1 {
				v1 = pr[k+1]
			} else {
				v1 = pr[pr[k+1]]
			}
			if OpCode(v)["B"] == 1 {
				v2 = pr[k+2]
			} else {
				v2 = pr[pr[k+2]]
			}
			if v1 < v2 {
				pr[pr[k+3]] = 1
			} else {
				pr[pr[k+3]] = 0
			}
			i += 3
		case 8:
			v1 := 0
			v2 := 0
			if OpCode(v)["C"] == 1 {
				v1 = pr[k+1]
			} else {
				v1 = pr[pr[k+1]]
			}
			if OpCode(v)["B"] == 1 {
				v2 = pr[k+2]
			} else {
				v2 = pr[pr[k+2]]
			}
			if v1 == v2 {
				pr[pr[k+3]] = 1
			} else {
				pr[pr[k+3]] = 0
			}
			i += 3
		case 99:
			return pr
		}
	}
	return pr
}

func readInput() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	in, err := strconv.Atoi(text)
	check(err)
	return in
}

func OpCode(x int) map[string]int {
	xs := strconv.Itoa(x)
	xm := strings.Split(xs, "")
	codeMap := make(map[string]int)
	if len(xm) >= 3 {
		codeMap["DE"], _ = strconv.Atoi(xs[len(xs) - 2:])
		if len(xs) - 3 < 0 {
			codeMap["C"] = 0
		} else {
			codeMap["C"], _ = strconv.Atoi(xs[len(xs) - 3:len(xs) - 2])
		}
		if len(xs) - 4 < 0 {
			codeMap["B"] = 0
		} else {
			codeMap["B"], _ = strconv.Atoi(xs[len(xs) - 4:len(xs) - 3])
		}
		if len(xs) - 5 < 0 {
			codeMap["A"] = 0
		} else {
			codeMap["A"], _ = strconv.Atoi(xs[len(xs) - 5:len(xs) - 4])
		}
	} else {
		codeMap["DE"] = x
	}
	return codeMap
}