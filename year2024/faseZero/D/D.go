package main

import (
	"fmt"
)

/*
Complexity O(1)
*/
func main() {
	var E int
	var V int
	_, _ = fmt.Scanf("%d %d", &E, &V)
	t := E / V
	minutos := E % V
	horas := (19 + t) % 24
	minutos = (60 * minutos) / V
	fmt.Printf("%02d:%02d\n",horas, minutos)
}