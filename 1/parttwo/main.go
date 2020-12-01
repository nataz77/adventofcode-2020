package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./input.txt")
	check(err)
	input := string(dat)
	split := strings.Split(input, "\n")

	for i := 0; i <= len(split); i++ {
		currentval, err := strconv.Atoi(split[i])
		check(err)
		for x := 0; x < len(split)-1; x++ {
			nextval, err := strconv.Atoi(split[x])
			check(err)
			if (currentval + nextval) == 2020 {
				fmt.Printf("Expense report match: %d", (currentval * nextval))
			}
		}
	}
}
