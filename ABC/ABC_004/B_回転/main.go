package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	var m [4][4]string
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			m[i][j] = read()
		}
	}
	for i := 3; i >= 0; i-- {
		fmt.Printf("%s %s %s %s\n", m[i][3], m[i][2], m[i][1], m[i][0])
	}
}

func read() string {
	sc.Scan()
	return sc.Text()
}
