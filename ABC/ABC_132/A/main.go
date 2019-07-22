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
	s := read()
	result := true
	for i := 1; i < len(s); i++ {
		if strings.Count(s, s[i-1:i]) != 2 {
			result = false
			break
		}
	}
	if result {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func read() string {
	sc.Scan()
	return sc.Text()
}
