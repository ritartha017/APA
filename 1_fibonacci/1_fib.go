package main

import (
        "fmt"
        "math"
        "os"
        "strconv"
)

func fib1(n int) int {
        if n < 2 {
                return n
        }
         return fib1(n-1) + fib1(n-2)
}

func fib2(n int) int {
        i := 1
        j := 0
        for k := 1; k <= n; k++ {
                j = i + j
                i = j - i
        }
        return j
}

func fib3(n int) int {
        var i, j, k, h, t int
        i, h = 1, 1
        j, k = 0, 0
        for n > 0 {
                if n%2 == 1 {
                        t = j * h
                        j = i*h + j*k + t
                        i = i*k + t
                }
                t = int(math.Pow(float64(h), 2))
                h = 2*k*h + t
                k = int(math.Pow(float64(k), 2)) + t
                n = n / 2
        }
        return j
}

func main() {
          arg := os.Args[1]
          val, _ := strconv.ParseInt(arg, 10, 0)
          fmt.Println(fib1(int(val)))
 //       fmt.Println(fib2(10))
 //       fmt.Println(fib3(10))
}

