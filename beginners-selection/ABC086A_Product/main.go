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
    a := readInt()
    b := readInt()
    s := GetEvenOdd(a*b)
    fmt.Printf("%s\n", s)
}

func readInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}

func IsEven(i int) bool {
    return i%2 == 0
}

func GetEvenOdd(i int) string {
    if IsEven(i) {
        return "Even"
    } else {
        return "Odd"
    }
}