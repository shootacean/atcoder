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
	t := readString()
	answer := solve(s, t)
	fmt.Println(answer)
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func solve(s, t string) (answer string) {
	answer = "You will lose"
	for i := 0; i < len([]rune(s)); i++ {
		_s := string([]rune(s)[i : i+1])
		_t := string([]rune(t)[i : i+1])
		if !match(_s, _t) {
			return
		}
	}
	answer = "You can win"
	return
}

func match(s, t string) bool {
	if s == "@" || t == "@" {
		if s == "@" && strings.Contains("atcoder@", t) {
			return true
		}
		if t == "@" && strings.Contains("atcoder@", s) {
			return true
		}
		return false
	} else if s != t {
		return false
	}
	return true
}
