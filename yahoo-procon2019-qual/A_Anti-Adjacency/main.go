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
	k := readInt()

	answer := solve(n, k)
	fmt.Printf("%v\n", answer)
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

func solve(n, k int) string {
	var count int
	for i := 1; i <= n; i++ {
		if !isEven(i) {
			count++
		}
	}
	if count >= k {
		return "YES"
	} else {
		return "NO"
	}
}

func isEven(n int) bool {
	return n%2 == 0
}