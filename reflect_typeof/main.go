package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
}

func main() {
	t := reflect.TypeOf(MyStruct{})

	fmt.Println(t.String()) // main.MyStruct
	fmt.Println(t.Kind())   // struct
	fmt.Println(t.Name())   // MyStruct
}
