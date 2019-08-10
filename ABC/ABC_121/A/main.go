// AtCode ABC 122 A
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
	H := readInt()
	W := readInt()
	h := readInt()
	w := readInt()
	all := H * W
	row := h * W
	col := w * (H - h)
	fmt.Println(all - (row + col))
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
