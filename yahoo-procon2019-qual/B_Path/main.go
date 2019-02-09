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

	a1 := readInt()
	b1 := readInt()
	a2 := readInt()
	b2 := readInt()
	a3 := readInt()
	b3 := readInt()

	answer := solve(a1, b1, a2, b2, a3, b3)
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

func solve(a1, b1, a2, b2, a3, b3 int) string {
	nums := []int{a1, b1, a2, b2, a3, b3}
	if sumSlice(nums) != 15 {
		return "NO"
	} else {
		return "YES"
	}
	//path1 := []int{a1, b1}
	//path2 := []int{a2, b2}
	//path3 := []int{a3, b3}
	//paths := [][]int{path1, path2, path3}
	//for i, p1 := range paths {
	//	a := p1[0]
	//	b := p1[1]
	//	otherPaths := append(paths[:i], paths[i+1:]...)
	//	for j, p2 := range otherPaths {
	//		if a == p2[0] || b == p2[0] || a == p2[1] || b == p2[1] {
	//			lastPath := append(otherPaths[:j], otherPaths[j+1:]...)
	//			if lastPath[0][0] == p2[0] || lastPath[0][1] == p2[0] || lastPath[0][0] == p2[1] || lastPath[0][1] == p2[1] {
	//				return "YES"
	//			}
	//		}
	//	}
	//}
	//return "NO"
}

func sumSlice(nums []int) (sum int) {
	for _, n := range nums {
		sum += n
	}
	return
}
