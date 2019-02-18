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
	m := readInt()

	var types []int
	for i := 0; i < m; i++ {
		types = append(types, 0)
	}
	for i := 0; i < n; i++ {
		// person
		k := readInt()
		for j := 0; j < k; j++ {
			// type
			x := readInt() - 1
			types[x] = types[x] + 1
		}
	}

	var answer int
	for i := 0; i < m; i++ {
		if types[i] == n {
			answer++
		}
	}
	fmt.Println(answer)
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
