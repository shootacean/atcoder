package main

import (
	"bufio"
	"fmt"
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

func solve(aList []int) (layer int) {
	deck := aList
	var currentSize int
	var choiceIndex, choiceNum int
	for i := 0; 1 <= len(deck); i++ {
		choiceIndex, choiceNum = maxIntSlice(deck)
		deck = append(deck[:choiceIndex], deck[choiceIndex+1:]...)
		if currentSize == 0 || choiceNum < currentSize {
			currentSize = choiceNum
			layer++
		}
	}
	return
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
