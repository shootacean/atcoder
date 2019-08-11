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
	x := readInt()
	if is753(x) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func is753(x int) bool {
	switch {
	case x == 3:
		return true
	case x == 5:
		return true
	case x == 7:
		return true
	default:
		return false
	}
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
