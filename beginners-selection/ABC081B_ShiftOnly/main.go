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
	//sc.Split(bufio.ScanWords)
	_ = readString()
	b := readString()

	ss := strings.Split(b, " ")
	var ns []int
	var n int
	var e error
	for _, s := range ss {
		n, e = strconv.Atoi(s)
		if e != nil {
			panic("Not int")
		}
		ns = append(ns, n)
	}

	answer := CountHalfLimit(ns)
	fmt.Printf("%v\n", answer)
}

// 標準入力から文字列を読み取る
func readString() string {
	sc.Scan()
	return sc.Text()
}

func CountHalfLimit(ns []int) int {
	if !IsEvenAll(ns) {
		return 0
	}
	count := 1
	_ns := ns
	var newN int
	for {
		for i, n := range _ns {
			newN = n / 2
			if newN == 1 || !IsEven(newN) {
				return count
			}
			_ns[i] = newN
		}
		count++
	}
	return count
}

// 偶数判定
func IsEven(i int) bool {
	return i%2 == 0
}

// 偶数判定
func IsEvenAll(ns []int) bool {
	for _, n := range ns {
		if !IsEven(n) {
			return false
		}
	}
	return true
}
