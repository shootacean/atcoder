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
	m := readInt()
	vv := solve(m)
	fmt.Println(vv)
}

func solve(m int) string {
	switch {
	case m < 100:
		return "00"
	case 100 <= m && m <= 5000:
		return fmt.Sprintf("%02d", m/100)
	case 6000 <= m && m <= 30000:
		return fmt.Sprintf("%02d", m/1000+50)
	case 35000 <= m && m <= 70000:
		return fmt.Sprintf("%02d", ((m/1000)-30)/5+80)
	case 70000 < m:
		return "89"
	default:
		return ""
	}
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
