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
	
	k := readInt()
	a := readInt()
	b := readInt()
	
	answer := solve(k, a, b)
	fmt.Printf("%v\n", answer)
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func readInt() int {
	i, e := strconv.Atoi(readString())
	if e != nil {
		panic(e)
	}
	return i
}

func solve(k, a, b int) int {
	bis := 1
	money := 0
	
	if b-a < 2 {
		// 換金する意味が無いので無条件でポケットを叩く
		return k + 1
	}
	for i := k; 0 < i; i-- {
		//fmt.Printf("Status: %d, %d\n", bis, money)
		// １円で交換する
		if money > 0 {
			//println("money to bis")
			money--
			bis += b
			continue
		}
		// １円と交換する
		if ( i > 1 && i < b-a) && bis >= a {
			//println("bis to money")
			bis -= a
			money++
			continue
		}
		// ポケットを叩く
		//println("pocket")
		bis++
	}
	//println("------")
	return bis
}
