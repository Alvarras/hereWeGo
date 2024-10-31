package dasar

import (
	"fmt"
	"testing"
)

func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(i)
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func TestContinue(t *testing.T) {
	test(10)
	test(20)
	test(30)
}
