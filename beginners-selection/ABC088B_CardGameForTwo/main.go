package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	
	n := readInt()
	var aList []int
	for i := 0; i < n; i++ {
		aList = append(aList, readInt())
	}
	answer := solve(aList)
	fmt.Printf("%d\n", answer)
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readInt() int {
	i, e := strconv.Atoi(readString())
	if e != nil {
		panic(e)
	}
	return i
}

func solve(aList []int) int {
	deck := aList
	var alice int
	var bob int
	var choiceIndex, choiceNum int
	for i := 0; 1 <= len(deck); i++ {
		choiceIndex, choiceNum = maxIntSlice(deck)
		deck = append(deck[:choiceIndex], deck[choiceIndex+1:]...)
		if isEven(i) {
			alice += choiceNum
		} else {
			bob += choiceNum
		}
	}
	return int(math.Abs(float64(alice - bob)))
}

func isEven(n int) bool {
	if n == 0 {
		return true
	}
	return n%2 == 0
}

func maxIntSlice(slice []int) (int, int) {
	var index, max int
	for i, n := range slice {
		if n > max {
			index = i
			max = n
		}
	}
	return index, max
}
