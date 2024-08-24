package main

import "fmt"

var count int


func swap(a , b int) (int , int) {
	return b , a
}
// count := 0;
func solve(n int, a string, b string, c string) {
	if n == 1 {
		fmt.Printf("Move %s to %s\n", a, c)
		count = count + 1
	} else {
		solve(n-1, a, c, b)
		solve(1, a, b, c)
		solve(n-1, b, a, c)
	}

}

func main() {
	fmt.Println("----------------------------------")
	count = 0
	var n_layer int
	fmt.Scanf("%d", &n_layer)
	solve(n_layer, "A", "B", "C")
	fmt.Println("count: ", count)
	fmt.Println("----------------------------------")

	a := 3;
	b := 5;
	a , b = swap(a, b)
	fmt.Printf("a: %d, b: %d\n" , a , b)
	fmt.Println("----------------------------------")

}
