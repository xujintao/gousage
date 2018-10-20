package main

import (
	"fmt"
	"reflect"
)

type user struct {
	Name string
	Age  int
}

func assignConfig(p interface{}) {
	pt := reflect.TypeOf(p).Elem()
	fmt.Println(pt)

	pv := reflect.ValueOf(p).Elem()
	fmt.Println(pv)

	// for i := 0; i < pt.NumField(); i++ {
	for i := 0; i < pv.NumField(); i++ {
		//name := pt.Field(i).Name
		pf := pv.Field(i)
		switch pf.Kind() {
		case reflect.String:
			pf.SetString("alice")
		case reflect.Int, reflect.Int64:
			pf.SetInt(20)
		default:
		}
	}
}

func main() {
	u := user{"bill", 24}
	fmt.Println(reflect.ValueOf(u))        //{bill 24}
	fmt.Println(reflect.ValueOf(u).Type()) //main.user
	//fmt.Println(reflect.ValueOf(u).Elem()) //panic: reflect: call of reflect.Value.Elem on struct Value
	fmt.Println(reflect.Indirect(reflect.ValueOf(u))) //{bill 24}

	fmt.Println("---")
	fmt.Println(reflect.ValueOf(&u))                   //&{bill 24}
	fmt.Println(reflect.ValueOf(&u).Type())            //*main.user
	fmt.Println(reflect.ValueOf(&u).Elem())            //{bill 24}
	fmt.Println(reflect.Indirect(reflect.ValueOf(&u))) //{bill 24}

	//u := user{"bill", 24}
	//assignConfig(&u)
	//fmt.Println(u)
}
