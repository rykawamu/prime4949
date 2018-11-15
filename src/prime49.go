package main

import (
    "fmt"
    "flag"
    "strconv"
    "strings"
)

type Prime4949 struct {
    PrimeNumbers [] int
    List49 [] string
}

func main() {
    var nCount = flag.Int("N", 9, "N=9")
    var nMax = flag.Int("MAX", 100, "MAX=100")
    var isMsgOpt = flag.Bool("v", false, "v=true")
    flag.Parse()
    if *nCount > *nMax {
        *nCount = * nMax
    }

    var pri Prime4949
    msg := pri.Execute(*nCount)
    if *isMsgOpt {
        msg = fmt.Sprintf("N=[%v] :", *nCount) + msg
    }
    fmt.Println(msg)
}

func (pri *Prime4949) Execute(nCount int) string {
    const firstPrimeNum = 2
    pri.PrimeNumbers = []int{firstPrimeNum}
    pri.List49 = []string{}
    for i := firstPrimeNum+1; len(pri.List49) < nCount; i += 2 {
       pri.addPrime(i)
    }
    msg := strings.Join(pri.List49 ,",")
    return msg
}

func (pri *Prime4949) addPrime(num int) {
    check := func () bool{
        if num < 2 {
            return false
        }
        isPrime := true
        for _, prime := range pri.PrimeNumbers {
            if num % prime == 0 {
                isPrime = false
                break
            }
        }
        return isPrime
    }
    if check() {
        pri.PrimeNumbers = append(pri.PrimeNumbers, num)
        if s := strconv.Itoa(num); strings.ContainsAny(s, "4 | 9") {
            pri.List49 = append(pri.List49, s)
        }
    }
}
