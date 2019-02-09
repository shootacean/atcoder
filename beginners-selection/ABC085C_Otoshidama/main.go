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
	
	N := readInt()
	Y := readInt()
	
	x, y, z := solve(N, Y)
	fmt.Printf("%d %d %d\n", x, y, z)
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

func solve(num, price int) (x, y, z int) {
	cannotSolve := func() (int, int, int) { return -1, -1, -1 }
	
	price = price / 1000
	
	limit10k := price / 10
	limit5k := price / 5
	
	for i10k := limit10k; 0 <= i10k; i10k-- {
		for i5k := limit5k; 0 <= i5k; i5k-- {
			if num < (i10k + i5k) {
				continue
			}
			i1k := num - i10k - i5k
			n := i10k + i5k + i1k
			p := i10k*10 + i5k*5 + i1k
			if num == n && price == p {
				x = i10k
				y = i5k
				z = i1k
				return
			}
		}
	}
	return cannotSolve()
}
