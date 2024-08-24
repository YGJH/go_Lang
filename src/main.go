package main

import "fmt"

var global_variable = 32

/* 整數
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64
uintptr

byte // uint8 別名
rune // int32 別名
*/

/*
float32 float64
*/

/* 複數
complex64 complex128
*/

/*
bool
*/

/*
string
*/

/* Printf , Println
" "雙引號 內可跳脫字元\t \n等

` `重音符 內則保留原始字串
*/

type Name struct {
	A string
	B bool
	C int
}

func main() {

	fmt.Println("--------------------------------------")
	fmt.Println("Hello World")
	fmt.Println("--------------------------------------")

	var a int
	fmt.Scanf("%d", &a)
	fmt.Println("--------------------------------------")
	fmt.Printf("a 的類型為： %T, a = %d\n", a, a) // a 的類型為 int
	var b float32
	b = 3.14
	fmt.Printf("b 的類型為： %T, b = %f\n", b, b)
	c := false
	fmt.Printf("c 的類型為： %T, c = %v\n", c, c)
	d := "Hello World"
	fmt.Printf("d 的類型為： %T, d = %s\n", d, d)
	fmt.Println("--------------------------------------")

	// //
	fmt.Println("--------------------------------------")
	fmt.Printf("%v	\n", Name{})
	fmt.Printf("%+v	\n", Name{})
	fmt.Printf("%#v	\n", Name{})
	fmt.Println("--------------------------------------")

	// SPrintf

	fmt.Println("--------------------------------------")
	s1 := "I"
	s2 := "am"
	s3 := "string"

	str1 := fmt.Sprintln(s1, s2, s3)
	fmt.Println(str1)

	str2 := fmt.Sprint(s1, s2, s3)
	fmt.Println(str2)

	fmt.Println("--------------------------------------")
	if a >= 3 {
		fmt.Println("a > 3")
	} else if a < -100 {
		fmt.Println("a < -100")
	} else {
		fmt.Println("!(a < -100) && !(a >= 3)")
	}
	fmt.Println("--------------------------------------")

	// y / m < x && x >  y / m2
	// y = mx
	// y = m2x + -2a
	// y + 2a / m2
	for i := a; i >= 0 ; i-- {
		f:=false
		for j := 0 ; j <= 2*a; j++ {
			if (i > j) || (j > (- i + 2*a)) {
				if f {
					break
				}
				fmt.Printf(" ")
			} else {
				f = true
				fmt.Printf("*")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("--------------------------------------")

}
