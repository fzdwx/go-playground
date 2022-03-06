package main

import (
	"fmt"
	"reflect"
)

func main() {

	var ff = 11.10

	t := reflect.TypeOf(ff)
	fmt.Println(t)

	v := reflect.ValueOf(&ff)
	fmt.Println(v.Elem())

	v.Elem().SetFloat(12.00)
	fmt.Println(ff)

	// customize type
	//ple := PeoPle{name: "like"}
	//t = reflect.TypeOf(ple)
	//fmt.Println(t)

	// https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/11.12.md
}
