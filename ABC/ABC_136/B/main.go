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
	c := 0
	for i := 1; i <= n; i++ {
		switch len(strconv.Itoa(i)) {
		case 1:
			c++
		case 3:
			c++
		case 5:
			c++
		}
	}
	fmt.Println(c)
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
