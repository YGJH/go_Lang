package main

// import "crypto/x509"
import "fmt"
// transform func(int) int 這代表函數的類型
func sum(a int , b int ,
	transform func(int) int) int {
		return transform(a) + transform(b)
}


func main() {
	// 這應該算 lambda 吧
	square := func(x int) int {
		return x * x
	}
	fmt.Println(sum(1 , 2 , square)) // 5
}
