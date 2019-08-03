// AtCode ABC 123 A
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
	a := readInt()
	_ = readInt()
	_ = readInt()
	_ = readInt()
	e := readInt()
	k := readInt()
	if e - a > k {
		fmt.Println(":(")
	} else {
		fmt.Println("Yay!")
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
