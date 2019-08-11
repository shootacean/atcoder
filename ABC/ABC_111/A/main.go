package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	n := read()
	answer := ""
	for _, r := range n {
		answer += convertBoth(string(r), "1", "9")
	}
	fmt.Println(answer)
}

func convertBoth(s, a, b string) string {
	switch s {
	case a:
		return b
	case b:
		return a
	default:
		return s
	}
}

func read() string {
	sc.Scan()
	return sc.Text()
}
