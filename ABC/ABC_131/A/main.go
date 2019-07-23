package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	s := read()
	prev := ""
	var result = ""
	for i := 1; i <= len(s); i++ {
		if prev == s[i-1:i] {
			result ="Bad"
			break
		}
		prev = s[i-1:i]
	}
	if result == "" {
		result = "Good"
	}
	fmt.Println(result)
}

func read() string {
	sc.Scan()
	return sc.Text()
}
