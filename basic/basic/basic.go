package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	fmt.Printf("%.3f\n",
		cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func sum(p int)  {
	var sum = 0
	for i := 0; i < p; i++ {
		sum += i
	}
	fmt.Println(sum)
}





//递归斐波那契数列 1, 1, 2, 3 ,5
func FbSumList(count int) int {
	if count == 1{
		return count
	}
	if count == 0{
		return 0
	}
	return FbSumList(count-1)+FbSumList(count-2)
}



//计算一个数的阶乘, 使用递归
//接受一个参数例如 5 = 1*2*3*4*5
func factorial(num int) int  {
	if num <= 1{
		return 1
	}
	return  num * factorial(num-1)
}


//结构体的概念
type person struct {
	name string
	city string
	age  int8
}


func main() {
	var sum =  factorial(10)
	fmt.Println(sum)
	//var i int
	//for i = 0; i < 10; i++ {
	//	fmt.Printf("%d\n", FbSumList(i))
	//sum(10)
	//fmt.Println("Hello world")
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
	//variableShorter()
	//fmt.Println(aa, ss, bb)
	//euler()
	//triangle()
	//consts()
	//enums()
}
