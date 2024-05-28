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
	N := 3
	M := 3
	K := 3
	L := 3
	K--
	L--
	quartos := [][]string{{"*","*","*"},{"*",".","*"},{"*","*","*"}}
	fmt.Printf("%d %d\n",N,M)
	fmt.Printf("%d %d\n",K,L)
	pontos := make([][]int,N)
	for i := range pontos {
		pontos[i] = make([]int,M) 
	}
	i := 0
	j := 0
	best_i := -1
	best_j := -1
	for i < N {
		j = 0
		for j < M {
			if quartos[i][j] == "*"{
				if K == 1{
					// pra esquerda
					if j - L >=0 {
						pontos[i][j-L]++
					}
					// pra direita
					if j + L < M {
						pontos[i][j+L]++
					}
					// pra cima
					if i - L >= 0 {
						pontos[i-L][j]++
					}
					// pra baixo
					if i + L < N {
						pontos[i+L][j]++
					}
				}else{
					// i = 0, j = 0
					// K = 2, L = 2
					// i = -2, j = -2
					// pra cima e esquerda
					if i - K >= 0 && j - L >=0{
						pontos[i-K][j-L]++
					}
					// i = 0, j = 0, N = 3, M = 3
					// K = 2, L = 2
					// i = -2, j = 2
					// pra cima e direita
					if i - K >= 0 && j + L < M{
						pontos[i-K][j+L]++
					}
					// i = 0, j = 0, N = 3, M = 3
					// K = 2, L = 2
					// i = 2, j = -2
					// pra baixo e esquerda
					if i + K < N && j - L >=0{
						pontos[i+K][j-L]++
					}
					// i = 0, j = 0, N = 3, M = 3
					// K = 2, L = 2
					// i = 2, j = 2
					// pra baixo e direita
					if i + K < N && j + L < M{
						pontos[i+K][j+L]++
					}
					if K != L {
						// inverter indices
						temp := K
						K = L
						L = temp
						// i = 0, j = 0
						// K = 2, L = 2
						// i = -2, j = -2
						// pra cima e esquerda
						if i - K >= 0 && j - L >=0{
							pontos[i-K][j-L]++
						}
						// i = 0, j = 0, N = 3, M = 3
						// K = 2, L = 2
						// i = -2, j = 2
						// pra cima e direita
						if i - K >= 0 && j + L < M{
							pontos[i-K][j+L]++
						}
						// i = 0, j = 0, N = 3, M = 3
						// K = 2, L = 2
						// i = 2, j = -2
						// pra baixo e esquerda
						if i + K < N && j - L >=0{
							pontos[i+K][j-L]++
						}
						// i = 0, j = 0, N = 3, M = 3
						// K = 2, L = 2
						// i = 2, j = 2
						// pra baixo e direita
						if i + K < N && j + L < M{
							pontos[i+K][j+L]++
						}
					}
				}
			}else if best_i == -1{
				best_i = i
				best_j = j
			}
			j++
		}
		i++
	}
	fmt.Printf("%d %d",best_i,best_j)
}
