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
	n := readInt()

	var monsters []int
	for i := 0; i < n; i++ {
		monsters = append(monsters, readInt())
	}

	fmt.Println(solve(n, monsters))
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

func solve(n int, monsters []int) int {
	slim := slimSlice(monsters)
	if len(slim) == 1 {
		return slim[0]
	}
	answer := gcd(slim[0], slim[1])
	for i, monster := range slim {
		if i <= 1 {
			continue
		}
		answer = gcd(answer, monster)
	}
	return answer
}

// 最大公約数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// sliceの重複要素を削除
func slimSlice(slice []int) []int {
	m := make(map[int]int)
	slim := make([]int, 0)
	for _, element := range slice {
		if _, ok := m[element]; !ok {
			m[element] = 0
			slim = append(slim, element)
		}
	}
	return slim
}

// 自力実装での最大公約数
func _solve(n int, monsters []int) int {

	slim := slimSlice(monsters)
	if len(slim) == 1 {
		return slim[0]
	}

	var max = 1
	var answer = 1
	var allOK = false
	//fmt.Println("--------")
	for i := 0; i < len(slim); i++ {
		allOK = true
		//fmt.Println("--------")
		fmt.Printf("i: %v\n", slim[i])
		for j := 0; j < len(slim); j++ {
			if i == j {
				continue
			}
			//fmt.Printf("j: %v\n", slim[j])
			if slim[j]%slim[i] != 0 {
				allOK = false
				break
			}
		}
		if allOK {
			// すべて割り切れたとき
			answer = slim[i]
			//fmt.Printf("answer: %v\n", answer)
			if max < answer {
				max = answer
			}
		}
	}
	return max
}
