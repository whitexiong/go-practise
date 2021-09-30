package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
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

func sum(p int) {
	var sum = 0
	for i := 0; i < p; i++ {
		sum += i
	}
	fmt.Println(sum)
}

//递归斐波那契数列 1, 1, 2, 3 ,5
func FbSumList(count int) int {
	if count == 1 {
		return count
	}
	if count == 0 {
		return 0
	}
	return FbSumList(count-1) + FbSumList(count-2)
}

//计算一个数的阶乘, 使用递归
//接受一个参数例如 5 = 1*2*3*4*5
func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * factorial(num-1)
}

//func Hnt(int count)  {
//
//}

//结构体的概念, 使用结构体创造一个人
type person struct {
	name    string
	city    string
	age     int8
	pointer *int
}

func bubbleSort(arr *[1000]int) {
	for i := range arr {
		for j := range arr {
			if arr[i] < arr[j] {
				swap(&arr[i], &arr[j])
			}
		}
	}
	fmt.Println(arr)
}

//交换数据
func swap(a *int, b *int) {
	*a, *b = *b, *a
}

//nlogin
func quickSort(arr *[1000]int, l int, r int) {
	//递归出口
	if l >= r {
		return
	}
	p := partition(arr, l, r)
	quickSort(arr, l, p-1) //左边
	quickSort(arr, p+1, r) //右边
}

func partition(arr *[1000]int, l int, r int) int {
	var v, j int
	v = arr[l] //选取最左边的索引
	j = l
	//遍历
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			swapArr(&arr[i], &arr[j+1])
			j++
		}
	}
	swapArr(&arr[l], &arr[j])
	return j
}

func swapArr(i *int, j *int) {
	*i, *j = *j, *i
}

func slice() {

}

//插入排序
//给定一个数组[4,2,1,3]
func insertSort(arr *[1000]int, len int) {
	for i := 1; i <= len-1; i++ {
		//待插入数字
		var temp int
		temp = arr[i]
		j := i - 1 //待比较索引
		//第一轮循环自己和自己比较 a,b   此时 a = j, 只有一个待排序数字
		//第二轮
		for ; j >= 0; j-- {
			if arr[j] > temp { // 4 > 5 ?
				arr[j+1] = arr[j] // 不停的覆盖原来的值，原地算法
			} else {
				break //跳出这个循环
			}
		}
		arr[j+1] = temp //第一轮循环完毕后 temp = arr[i] = 5, arr[j+1] = 5 而 arr[j+1] = arr[i]
	}
}

func selectSort(arr *[1000]int, len int) {
	for i := 0; i < len; i++ {
		var min, index int
		var j = i
		min = arr[j]
		for ; j < len; j++ {
			if arr[j] < min {
				min = arr[j]
				index = j
			}
		}
		//不交换位置
		if index == 0 {
			continue
		} else {
			swap(&arr[i], &arr[index])
		}
	}
}

//归并 srot
func mergeSort(arr *[10]int) {

}

//randNumCount
func randNumber(randNumCount int, ranges int) [1000]int {
	var randArr [1000]int
	for i := 0; i < randNumCount; i++ {
		randArr[i] = rand.Intn(ranges)
	}
	return randArr
}

func main() {
	var arr [1000]int
	arr = randNumber(1000, 1000)
	start := time.Now() // 获取当前时间
	//var arr = [...]int{4,5,10,1,220,55,785,999,2,0}
	//quickSort(&arr, 0, len(arr)-1)
	insertSort(&arr, len(arr))
	//bubbleSort(&arr)
	//selectSort(&arr,len(arr))
	elapsed := time.Now().Sub(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
	fmt.Println(arr)
	//var a,b = 1,2
	//swap(&a, &b)
	//fmt.Println(a)
	//var arr = [4] int{3,2,1,4}
	//bubbleSort(&arr)
	//var p1 person
	//var a = 10;
	//p1.age = 23
	//p1.pointer = &a
	//fmt.Println("%v", p1.pointer)
	//var sum =  factorial(10)
	//fmt.Println(sum)
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
