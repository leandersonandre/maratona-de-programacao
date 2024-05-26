package main

import (
	"fmt"
    "sort"
)
type node struct{
	value int 
	next *node
}

type Stack struct{
	top *node
	size int
}
func New() *Stack{
	l := Stack{}
	return &l
}

func (s *Stack) Size() int{
	return s.size
}

func (s *Stack) IsEmpty() bool{
	return s.size == 0
}

func (s *Stack) Push(value int){
	newNode := node{value,s.top}
	s.top = &newNode
	s.size++
}


func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return  0, fmt.Errorf("stack is empty")
	}
	return s.top.value, nil
}


func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return  0, fmt.Errorf("stack is empty")
	}
	old := s.top
	s.top = s.top.next
	s.size--
	return old.value, nil
}

func main() {
	var N int // comunidades
	var M int // qtde de trechos
	var K int // qtde de festas
	
	_, _ = fmt.Scanf("%d %d %d",&N, &M, &K)
	entradas := make([][3]int, M)
	for i := range entradas{
		_, _ = fmt.Scanf("%d %d %d", &entradas[i][0], &entradas[i][1], &entradas[i][2])
	}
	
	type caminho struct{
		destino int 
		peso int
	}
	grafo := make([][]caminho, N)
	//entradas := [][]int { {1, 4, 0}, {3, 4, 0}, {4 ,5 ,0}, {3, 5, 0},{1, 5, 0}, {3, 2, 2}, {5, 2, 2}, {1, 2, 2}}
	//entradas := [][]int { {1, 4, 0}, {3, 4, 3}, {4 ,5 ,2}, {3, 5, 8}, {1, 5, 5}, {3, 2, 9}, {5, 2, 8}, {1, 2, 2}}
	for _, v := range entradas {
		c := caminho{v[1]-1,v[2]}
		grafo[v[0]-1] = append(grafo[v[0]-1], c)
		// caminho inverso
		c = caminho{v[0]-1,v[2]}
		grafo[v[1]-1] = append(grafo[v[1]-1], c)
	}
	// Depth First Search
	// Time complexity: O(V + E)
	// V - vertices
	// E - arestas/ edges
	visitado := make([]bool, N)
	peso := make([]int, N)
	// Primeira cidade visitada
	// Saída a partir da cidade 1
	visitado[0] = true
	peso[0] = 0
	// Stack : O(1) para todos metodos
	pilhaDestino := New()
	pilhaPeso := New()
	for _, c := range grafo[0]{
		pilhaDestino.Push(c.destino)
		pilhaPeso.Push(c.peso)
	}
	for !pilhaDestino.IsEmpty() {
		destino,_ := pilhaDestino.Pop()
		pesoCaminho,_ := pilhaPeso.Pop()
		if(!visitado[destino]){
			visitado[destino] = true
			peso[destino] = pesoCaminho
			for _, c := range grafo[destino]{
				pilhaDestino.Push(c.destino)
				pilhaPeso.Push(c.peso)
			}
		}else if peso[destino] > pesoCaminho{
				peso[destino] = pesoCaminho
		}
	}
	// Ordenação : O(nlogn)
	sort.Slice(peso, func(i, j int) bool {
		return peso[i] < peso[j]
	})
	fmt.Printf("%d\n",peso[K-1])
}