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
	fmt.Println(n * 2)
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
