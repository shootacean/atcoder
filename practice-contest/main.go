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
    c := readInt()
    var s string
    if sc.Scan() {
        s = sc.Text()
    }
    fmt.Printf("%d %s\n",a+b+c , s)
}

func readInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    if e != nil {
        panic(e)
    }
    return i
}