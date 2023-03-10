package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 return false
	for i:= 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i === 0 return false
	}
	return true
}

func main(){
	n := 10
	if isPrime(n){
		fmt.Println("%d is a prime number\n", n)
	} else {
		fmt.Println("%d is not a prime number\n", n)
	}
}
