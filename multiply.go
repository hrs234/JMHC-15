package main

import (
	"fmt"
	"strings"
	"strconv"
)

func multi(data string) int {
	// split the data into 2 section
	split := strings.Split(data, " ")
	
	// 0: many bulbs
	manyBulb, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Println("ERROR: ", err)
		// return []int{}
	}
	
	// 1: many switchs turned
	manySwitch, err := strconv.Atoi(split[1])
	if err != nil {
		fmt.Println("ERROR: ", err)
		// return []int{}
	}
	
	arr := []int{}
	// generate the bulb
	for i := 0; i < manyBulb; i++ {
		arr = append(arr, 0)
	}

	// step by step multiplication edit 
	// if the bulbs is on then off
	for b := 1; b <= manySwitch; b++ {
		for c := 0; c < manyBulb; c+=b {
			if arr[c] == 0 {
				arr[c] = 1
			} else {
				arr[c] = 0
			}
		}
	}

	// last: count all the array if their value is 1 which mean on
	count := 0
	for d := 0; d < len(arr); d++ {
		if arr[d] == 1 {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println("h#1: ", multi("4 1"))

	fmt.Println("h#2: ", multi("10 3"))
	fmt.Println("h#3: ", multi("10 4"))
	// multi("10 1")
}