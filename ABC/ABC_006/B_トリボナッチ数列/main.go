package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

//var memo = make(map[int]int)

func main() {
	sc.Split(bufio.ScanWords)
	n := readInt()
	a := solve(int(n))
	fmt.Println(a)
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

func solve(n int) int {
	return tori(n)
}

// 参考：http://topcoder.g.hatena.ne.jp/nodchip/20140405/1396707671
func tori(n int) int {
	a := 0
	b := 0
	c := 1
	for i:= 2; i <= n; i++ {
		d := (a + b + c) % 10007
		a = b
		b = c
		c = d
	}
	return a
}

// 解けない(桁あふれ？)
//func tori_memo(n int) int {
//	m, exists := memo[n]
//	if exists {
//		return m
//	}
//	var t int
//	if n <= 0 || n == 1 || n == 2 {
//		return 0
//	}
//	if n == 3 {
//		return 1
//	}
//	var i int
//	for i = 1; i <= 3; i++ {
//		t += tori(n - i)
//	}
//	memo[n] = t
//	return t
//}
