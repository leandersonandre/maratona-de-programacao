package main

import (
	"fmt"
)

/*
Complexity O(N)
*/
func main() {
	/*
	N := 5
	quartos := []string {"Alice","Bob","Clara", "Dora", "Eve"}
	idades := []string {"Bob","Clara","Eve", "Dora", "Alice"}
	*/
	var N int
	_, _ = fmt.Scanf("%d", &N)
	quartos := make([]string, N)
	for i := range quartos{
		_, _ = fmt.Scanf("%s", &quartos[i])
	}
	idades := make([]string, N)
	for i := range quartos{
		_, _ = fmt.Scanf("%s", &idades[i])
	}
	i := 0
	j := 0
	for i < N {
		fmt.Printf("%s",idades[j])
		if idades[j] == quartos[i] {
			j++
		}
		i++
		if i < N {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
