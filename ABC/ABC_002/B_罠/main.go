package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	s := readString()
	answer := s
	for _, x := range []string{"a", "i", "u", "e", "o"} {
		answer = strings.Replace(answer, x, "", -1)
	}
	fmt.Println(answer)
}

func readString() string {
	sc.Scan()
	return sc.Text()
}
