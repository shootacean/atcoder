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
	var p int
	c := 0
	for i := 1; i <= n; i++ {
		p = readInt()
		if p != i {
			c++
		}
		if c > 2 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
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
