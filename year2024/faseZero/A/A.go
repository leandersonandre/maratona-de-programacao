package main
 
import (
				"fmt"
)
	

/*
Complexity O(1)
*/
func main() {
	var a int
	var b int
	var c int
	_, _ = fmt.Scanf("%d %d %d", &a, &b, &c)
	d := 6 - a - b- c
	fmt.Printf("%d\n",d)
}
