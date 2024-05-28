package main

import (
	"fmt"
    "math"
)

/*
Complexity O(N)
1 is not a primer
*/
func main() {
	/*
	T := 4
	linhas := []int {3, 15, 5, 21}
	fmt.Printf("%d\n",T)
	fmt.Printf("%v\n",linhas)
	*/
	var T int
	_, _ = fmt.Scanf("%d", &T)
	linhas := make([]int, T)
	for i := range linhas{
		_, _ = fmt.Scanf("%d", &linhas[i])
	}
	checked := make([]bool, 1_000_000_000)
	isPrime := make([]bool, 1_000_000_000)
	//O(T)
	for _,v := range linhas{
		BEST_X := -1
		BEST_Y := -1
		X := 1
		Y := v - 1
		//O(v/2)
		for X <= v / 2 {
			if !checked[X] {
				checked[X] = true
				isPrime[X] = isNumberPrime(X)
			}
			if !checked[Y] {
				checked[Y] = true
				isPrime[Y] = isNumberPrime(Y)
			}
			if !isPrime[X] && !isPrime[Y]{
				if BEST_X == -1 ||
					math.Abs(float64(X -Y)) < math.Abs(float64(BEST_X - BEST_Y)) {
					BEST_X = X
					BEST_Y = Y
				}
			}
			X++
			Y--
		}
		if(BEST_X > -1){
			fmt.Printf("%d %d\n", BEST_X, BEST_Y)
		}else{
			fmt.Printf("-1\n")
		}
	}
}


//Complexity O(sqrt(N))
func isNumberPrime(n int) bool{
	if n <= 1 {
		return false
	}
	// Check if n=2 or n=3 
	if n == 2 || n == 3 {
		return true
	}
	// Check whether n is divisible by 2 or 3 
	if n % 2 == 0 || n % 3 == 0 {
		return false
	}

	// Check from 5 to square root of n 
	// Iterate i by (i+6)
	i := 5.0
	for i <= math.Sqrt(float64(n)) {
		if n % int(i) == 0 || n % int(int(i) + 2) == 0{
			return false
		}
		i = i + 6
	}
	return true
}