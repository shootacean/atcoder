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
	n := float64(readInt())
	var answer float64
	for i := 1.0; i <= n; i++ {
		answer += (i*10000) * (1/n)
	}
	fmt.Println(int(answer))
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
