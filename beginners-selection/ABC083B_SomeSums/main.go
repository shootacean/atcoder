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
	a := readInt()
	b := readInt()

	answer := solve(n, a, b)
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

func solve(n, a, b int) (sum int) {
	var numbers []int
	for i := 1; i < n+1; i++ {
		some := sumSlice(strconv.Itoa(i))
		if a <= some && some <= b {
			numbers = append(numbers, i)
		}
	}
	for _, number := range numbers {
		sum += number
	}
	return
}

func sumSlice(ss string) (sum int) {
	for i := 0; i < len([]rune(ss)); i++ {
		n, _ := strconv.Atoi(string([]rune(ss)[i : i+1]))
		sum += n
	}
	return
}
