package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.IsPrime(4))
	fmt.Println(piscine.IsPrime(5))
	fmt.Println(piscine.IsPrime(13))
	fmt.Println(piscine.IsPrime(16))
	fmt.Println(piscine.IsPrime(17))
	fmt.Println(piscine.IsPrime(1))
	fmt.Println(piscine.IsPrime(0))
	fmt.Println(piscine.IsPrime(-1))
}
