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
	d := readInt()
	answer := solve(float64(n), float64(d))
	fmt.Println(answer)
}

func solve(n, d float64) float64 {
	return math.Ceil(n / ((d * 2) + 1))
}

func read() string {
	sc.Scan()
	return sc.Text()
}

func readInt() int {
	i, e := strconv.Atoi(read())
	if e != nil {
		panic(e)
	}
	return i
}
