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

	coin500 := readInt()
	coin100 := readInt()
	coin50 := readInt()
	pay := readInt()

	answer := solveCoins(coin500, coin100, coin50, pay)
	fmt.Printf("%v\n", answer)
}

// 標準入力から文字列を読み取る
func readString() string {
	sc.Scan()
	return sc.Text()
}

// 標準入力を数値として読み取る
func readInt() int {
	i, e := strconv.Atoi(readString())
	if e != nil {
		panic(e)
	}
	return i
}

func solveCoins(coin500, coin100, coin50, pay int) int {
	count := 0
	var price int
	for i500 := 0; i500 <= coin500; i500++ {
		if pay < (500 * i500) {
			break
		}
		for i100 := 0; i100 <= coin100; i100++ {
			if pay < (500*i500)+(100*i100) {
				break
			}
			for i50 := 0; i50 <= coin50; i50++ {
				price = (500 * i500) + (100 * i100) + (50 * i50)
				if pay < price {
					break
				}
				if pay == price {
					count++
				}

			}
		}
	}
	return count
}
