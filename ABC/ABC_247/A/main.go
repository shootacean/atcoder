package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	s := read()
	var a []string
	a = append(a, "0")
	for i, _s := range strings.Split(s, "") {
		if i == i - 1 {
			break;
		}
		a = append(a, _s)
		fmt.Println(_s)
	}
	fmt.Println(strings.Join(a, ""))
}


func read() string {
	sc.Scan()
	return sc.Text()
}
