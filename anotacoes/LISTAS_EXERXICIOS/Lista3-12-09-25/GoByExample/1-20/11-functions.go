package main

import "fmt"

func soma(a int, b int) int {

    return a + b
}

func somaSoma(a, b, c int) int {
    return a + b + c
}

func main() {

    res := soma(1, 2)
    fmt.Println("1+2 =", res)

    res = somaSoma(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}