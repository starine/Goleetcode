package main

import (
	"fmt"
	"unsafe"
)

type slice struct {
	array unsafe.Pointer //切片底层数组的起始位置
	len   int            //切片长度
	cap   int            //切片容量
}

func main() {
	var s []int            //变量声明 nil切片，不需要分配内存
	s1 := []int{}          //长度为0的切片，空切片，指长度为空，其值并不是nil
	s2 := []int{1, 2, 3}   //长度为3的切片
	fmt.Println(s)         //[]
	fmt.Println(s1)        //[]
	fmt.Println(s2)        //[1 2 3]
	fmt.Println(s == nil)  //true
	fmt.Println(s1 == nil) //false

	s3 := make([]int, 12)
	s4 := make([]int, 10, 100)
	fmt.Println(s3) //[0 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(s4) //[0 0 0 0 0 0 0 0 0 0]

	array := [5]int{1, 2, 3, 4, 5}
	s5 := array[0:2] //从数组中切取
	s6 := s5[0:1]    //从切片中切取
	fmt.Println(s5)  //[1 2]
	fmt.Println(s6)  //[1]

	slice := *new([]int)
	fmt.Println(slice)        //[]
	fmt.Println(slice == nil) //true

	s7 := make([]int, 0)
	fmt.Printf("slice:%v len:%d cap:%d\n", s7, len(s7), cap(s7)) //slice:[] len:0 cap:0

	s7 = append(s7, 1)
	fmt.Printf("slice:%v len:%d cap:%d\n", s7, len(s7), cap(s7)) //slice:[1] len:1 cap:1

	s7 = append(s7, 2, 3, 4)
	fmt.Printf("slice:%v len:%d cap:%d\n", s7, len(s7), cap(s7)) //slice:[1 2 3 4] len:4 cap:4

	s7 = append(s7, []int{5, 6}...)                              //当append前，切片空间不够时，会先创建2倍的cap大小，再拷贝返回新的切片
	fmt.Printf("slice:%v len:%d cap:%d\n", s7, len(s7), cap(s7)) //slice:[1 2 3 4 5 6] len:6 cap:8

	s8 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s9 := make([]int, 5)
	copy(s9, s8)
	fmt.Printf("slice:%v len:%d cap:%d\n", s8, len(s8), cap(s8)) //slice:[1 2 3 4 5 6 7 8 9] len:9 cap:9
	fmt.Printf("slice:%v len:%d cap:%d\n", s9, len(s9), cap(s9)) //slice:[1 2 3 4 5] len:5 cap:5

	a := [3]int{1, 2, 3}
	b := a[1:2]                     //简单表达式
	b = append(b, 4)                //此时元素a[2]将由3变为4
	fmt.Printf("a:%v b:%v\n", a, b) //a:[1 2 4] b:[2 4]

	a1 := [3]int{1, 2, 3}
	b1 := a1[1:2:2]                     //扩展表达式
	b1 = append(b1, 4)                  //此时元素a[2]将由3变为4
	fmt.Printf("a1:%v b1:%v\n", a1, b1) //a1:[1 2 3] b1:[2 4]

	SliceCap()
	SlicePrint()
	SliceExtend()
	SliceExpress()
}
func SliceCap() {
	var array [10]int
	var slice = array[5:6]
	fmt.Printf("len(slice)=%d\n", len(slice))
	fmt.Printf("cap(slice)=%d\n", cap(slice))
}
func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}
func SlicePrint() {
	s1 := make([]int, 2, 5)
	s1 = []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	fmt.Printf("s1:%v len:%d cap:%d\n", s1, len(s1), cap(s1)) //slice:[1 2 3 4 5 6 7 8 9] len:9 cap:9
	fmt.Printf("s2:%v len:%d cap:%d\n", s2, len(s2), cap(s2)) //slice:[1 2 3 4 5] len:5 cap:5
	SliceRise(s1)
	SliceRise(s2)
	fmt.Println(s1, s2)
}
func SliceExtend() {
	s := make([]int, 0, 10)
	s1 := append(s, 1, 2, 3)
	s2 := append(s1, 4)
	s2[0] = 5
	fmt.Println(&s1[0] == &s2[0])
	fmt.Printf("s:%v len:%d cap:%d\n", s, len(s), cap(s))
	fmt.Printf("s1:%v len:%d cap:%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%v len:%d cap:%d\n", s2, len(s2), cap(s2))
}
func SliceExpress() {
	orderLen := 5
	order := make([]uint16, 2*orderLen)

	pollorder := order[:orderLen:orderLen]
	lockorder := order[orderLen:][:orderLen:orderLen] //连续切片

	fmt.Printf("len(pollorder)=%d\n", len(pollorder))
	fmt.Printf("cap(pollorder)=%d\n", cap(pollorder))
	fmt.Printf("len(lockorder)=%d\n", len(lockorder))
	fmt.Printf("cap(lockorder)=%d\n", cap(lockorder))
}
