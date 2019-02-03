package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
    sc.Split(bufio.ScanWords)
    s := readString()
    answer := strings.Count(s, "1")
    fmt.Printf("%v\n", answer)
}

func readString() string {
    sc.Scan()
    return sc.Text()
}