package main

import (
	"fmt"
	"math/big"

	"github.com/razor-87/bprimes"
)

func main() {
	numbers := []uint{2, 3, 5, 7, 900001, 1000000, 13756265695458089029, 13756265695458089031, 13496181268022124907}
	limit := bprimes.Limit()
	fmt.Printf("limit %d\n", limit)
	for _, n := range numbers {
		if n <= limit {
			fmt.Printf("IsPrime checked %d: %t\n", n, bprimes.IsPrime(n))
		} else {
			fmt.Printf("ProbablyPrime checked %d: %t\n", n, new(big.Int).SetUint64(uint64(n)).ProbablyPrime(20))
		}
	}
}
