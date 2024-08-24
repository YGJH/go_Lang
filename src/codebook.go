
package main
 
import (
	"bufio"
	"bytes"
	"math"
	"os"
	//"sort"
	//"rand"
	"strconv"
	"strings"
)
 
const (
	io = "std"
)
 
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inputFile, _ := os.Open("input.txt")
	defer inputFile.Close()
	if io == "file" {
		scanner = bufio.NewScanner(inputFile)
	}
	const nCapacity = 4
	const maxCapacity = nCapacity * 1024 * 1024 // nCapacity Mi
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	writer := bufio.NewWriter(os.Stdout)
	outputFile, _ := os.Open("output.txt")
	defer outputFile.Close()
	if io == "file" {
		writer = bufio.NewWriter(outputFile)
	}
	defer writer.Flush()
	scanner.Scan()
	t := atoi(scanner.Text())
	for ; t > 0; t-- {
		scanner.Scan()
		n := atoi(scanner.Text())
		if n%2 == 0 {
			writer.WriteString("-1\n")
			continue
		}
		var a []int
		for i := 1; i <= n; i++ {
			a = append(a, i)
		}
		for i := 1; i < n; i += 2 {
			a[i], a[i+1] = a[i+1], a[i]
		}
		for _, v := range a {
			writer.WriteString(itoa(v))
			writer.WriteString(" ")
		}
		writer.WriteString("\n")
	}
}
 
func FastConcat(sl []string) string {
	var b bytes.Buffer
	for _, v := range sl {
		b.WriteString(v)
	}
	return b.String()
}
 
func slatoi(sl []string) (r []int) {
	for _, v := range sl {
		r = append(r, atoi(v))
	}
	return
}
 
func pow2(n int) int {
	return powint(2, n)
}
 
func isPow2(x int) bool {
	return x > 0 && (x&(x-1)) == 0
}
 
func powint(n int, m int) int {
	return int(math.Pow(float64(n), float64(m)))
}
 
func splite(s string) []string {
	return split(s, "")
}
 
func splits(s string) []string {
	return split(s, " ")
}
 
func split(s string, c string) (sl []string) {
	sl = strings.Split(s, c)
	return
}
 
func atoi(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}
 
func itoa(i int) (s string) {
	s = strconv.Itoa(i)
	return
}
 
func ftoa(n float64, t int) string {
	return strconv.FormatFloat(n, 'f', t, 64)
}
 
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
 
func isPalindrome(s string) bool {
	return reverse(s) == s
}
 
func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}
 
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
 
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
 
func isPrime(n int) bool {
	if n < 2 || n%2 == 0 {
		return false
	}
	if n < 4 {
		return true
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
 
func FactRec(n int) int {
	if n == 1 {
		return 1
	}
	return n * FactRec(n-1)
}
 
func FactIt(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
 
func GCD(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
 
func LCM(a int, b int) int {
	return (a / GCD(a, b)) * b
}
 
func PrimeFactors(n int) (map[int]int, int) {
	p := make(map[int]int)
	f := 0
	for n%2 == 0 {
		p[2]++
		n = n / 2
		f++
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			p[i]++
			n = n / i
			f++
		}
	}
	if n > 2 {
		p[n]++
		f++
	}
	return p, f
}
 
func com(m int, n int) (c int) {
	// сочетания без повторений: n -- всего, m -- нужно выбрать
	c = FactRec(n) / FactRec(m) / FactRec(n-m)
	return
}
 
func comwr(m int, n int) (c int) {
	// сочетания с повторенями: n -- всего, m -- нужно выбрать
	c = FactRec(n+m-1) / FactRec(m) / FactRec(n-1)
	return
}
 
func per(m int, n int) (a int) {
	// размещения без повторений: n -- всего нужно разместить, m -- всего мест
	a = FactRec(n) / FactRec(n-m)
	return
}
 
func perwr(m int, n int) (a int) {
	//размещения с повторениями: n -- всего нужно разместить, m -- всего мест
	a = powint(n, m)
	return
}
 
func arr(n int) (p int) {
	// перестановки без повторений: n -- всего нужно разметить по n местам
	p = FactRec(n)
	return
}
 
func arrwr(nk []int, n int) (p int) {
	// перестановки с повторениями: nk -- количество каждого из k типов различных предметов
	p = FactRec(n)
	for _, v := range nk {
		p /= FactRec(v)
	}
	return
}
 
/*
 
	func bubbleSort(a []int) []int { // O(la^2)
		la := len(a)
		if la < 2 {
			return a
		}
		for i := 0; i < la; i++ {
			for j := 0; j < la-i-1; j++ {
				if a[j] > a[j+1] {
					a[j], a[j+1] = a[j+1], a[j]
				}
			}
		}
		return a
	}
 
	func quickSort(a []int) []int { // O(la^2)
		la := len(a)
		if la < 2 {
			return a
		}
		l, r := 0, la-1
		p := rand.Int() % la
		a[p], a[r] = a[r], a[p]
		for i := range a {
			if a[i] < a[r] {
				a[i], a[l] = a[l], a[i]
				l++
			}
		}
		return a
	}
 
	func insetSort(a []int) []int { // O(la^2)
		la := len(a)
		if la < 2 {
			return a
		}
		i := 1
		for i < la {
			j := i
			for j >= 1 && a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				j--
			}
			i++
		}
		return a
	}
 
	func selectSort(a []int) []int { // O(la^2)
		la := len(a)
		if la < 2 {
			return a
		}
		i := 1
		for i < la-1 {
			j := i + 1
			m := i
			if j < la {
				if a[j] < a[m] {
					m = j
				}
				j++
			}
			if m != i {
				a[i], a[m] = a[m], a[i]
			}
			i++
		}
		return a
	}
 
 
	func mergeFunc(fp []int, sp []int) []int {
		n := make([]int, len(fp)+len(sp))
		fpi := 0
		spi := 0
		ni := 0
		for fpi < len(fp) && spi < len(sp) {
			if fp[fpi] < sp[spi] {
				n[ni] = fp[fpi]
				fpi++
			} else if sp[spi] < fp[fpi] {
				n[ni] = sp[spi]
				spi++
			}
			ni++
		}
		for fpi < len(fp) {
			n[ni] = fp[fpi]
			fpi++
			ni++
		}
		for spi < len(sp) {
			n[ni] = sp[spi]
			spi++
			ni++
		}
		return n
	}
 
	func mergeSort(a []int) []int { // O(la*log(la))
		la := len(a)
		if la < 2 {
			return a
		}
		fp := mergeSort(a[0 : la/2])
		sp := mergeSort(a[la/2:])
		return mergeFunc(fp, sp)
	}
 
	func countSort(a []int) []int { // O(la+k)
		la := len(a)
		if la < 2 {
			return a
		}
		m:= a[0]
		i := 1
		for i < la {
			if a[i] > m {
				m = a[i]
			}
			i++
		}
		b:= make([]int, m + 1)
		j := 0
		for j < la {
			b[a[j]]++
			j++
		}
		lb:=len(b)
		k := 1
		for k < lb {
			b[k] += b[k - 1]
			k++
		}
		r := make([]int, la)
		m = 0
		for m < la {
			r[b[a[m]] - 1] = a[m]
			b[a[m]]--
			m++
		}
		return r
	}
 
*/