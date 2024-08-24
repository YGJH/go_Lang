package main

import(
	"fmt"
)

func main() {
	var T int
	fmt.Scanf("%d\n" , &T)
	// fmt.Printf("T = %d\n" , T)
	for T>0 {
		T = T - 1
		var n int
		fmt.Scanf("%d\n" , &n)
		// fmt.Printf("n = %d\n" , n)
		var tmp int
		count:=0
		mp := make(map[int]int)
		for i:=0 ; i < n ; i++{
			fmt.Scanf("%d", &tmp)
			// fmt.Printf("tmp = %d\n"  , tmp)
			mp[tmp]+=1;
			count = max(mp[tmp] , count)
			if i == n-1 {
				fmt.Scanf("\n")
			}

		}
		fmt.Printf("%d\n" , (n - count))
	}
}