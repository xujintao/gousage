package main

import (
	"fmt"
	"unsafe"
)

func main() {
	n1 := unsafe.Sizeof(true)                //1个字节
	n2 := unsafe.Sizeof(123)                 //8个字节
	n31 := unsafe.Sizeof(float32(3.1))       //4个字节
	n32 := unsafe.Sizeof(float64(3.1))       //8个字节
	n41 := unsafe.Sizeof(complex(4.1, 4.11)) //16个字节
	// n42 := unsafe.Sizeof(complex64(4.1, 4.11))
	n5 := unsafe.Sizeof("123") //16个字节
	var t6 *string
	n6 := unsafe.Sizeof(t6)          //8个字节
	n71 := unsafe.Sizeof(struct{}{}) //0个字节
	n72 := unsafe.Sizeof(struct {    //32个字节
		a bool   //8个字节
		b int    //8个字节
		c string //16个字节
	}{})
	n73 := unsafe.Sizeof(struct { //32个字节
		c string //16个字节
		a bool   //8个字节
		b int    //8个字节
	}{})
	n74 := unsafe.Sizeof(struct { //32个字节
		c string //16个字节
		b int    //8个字节
		a bool   //8个字节
	}{})
	n75 := unsafe.Sizeof(struct { //24个字节
		c string //16个字节
		b int32  //4个字节
		a bool   //4个字节
	}{})
	//能把n7系列看懂的都是人才
	n8 := unsafe.Sizeof([]int{1, 2, 3})                     //24个字节
	n9 := unsafe.Sizeof(map[string]int{"123": 1, "456": 2}) //8个字节
	n10 := unsafe.Sizeof(func(int, string) {})              //8个字节
	n11 := unsafe.Sizeof(interface{}(1))                    //16个字节
	c := make([]int64, 0)
	n12 := unsafe.Sizeof(c) //24个字节

	fmt.Println(n1, n2, n31, n32, n41, n5, n6, n71, n72, n73, n74, n75, n8, n9, n10, n11, n12)

	a := [3]string{"1", "2", "3"}
	b := [3]int{1, 2, 3}
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b))
}
