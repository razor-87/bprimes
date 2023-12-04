package main

import (
	"fmt"
	"math/big"

	"github.com/razor-87/bprimes"
)

func main() {
	n := new(big.Int)
	for i := 0; i <= bprimes.Limit(); i++ {
		if bprimes.IsPrime(uint32(i)) != n.SetInt64(int64(i)).ProbablyPrime(10) {
			panic(i)
		}
		if i%1_000_000 == 0 {
			fmt.Println("done:", i)
		}
	}
}
