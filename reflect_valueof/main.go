package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	A string `json:"a"`
}

func main() {
	v := reflect.ValueOf(MyStruct{A: "Foo"})

	fmt.Println(v.String())    // <main.MyStruct Value>
	fmt.Println(v.Interface()) // {Foo}
	fmt.Println(v.Kind())      // struct

	if v.Kind() == reflect.Struct {
		fmt.Println(v.Field(0).Interface()) // Foo
		if v.Field(0).CanSet() {
			v.Field(0).Set(reflect.ValueOf("Bar"))
		}
		fmt.Println(v.Field(0).Interface()) // Foo
	}

	v = reflect.ValueOf(&MyStruct{A: "Foo"})

	fmt.Println(v.String())    // <main.MyStruct Value>
	fmt.Println(v.Interface()) // {Foo}
	fmt.Println(v.Kind())      // ptr

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() == reflect.Struct {
		fmt.Println(v.Field(0).Interface()) // Foo
		if v.Field(0).CanSet() {
			v.Field(0).Set(reflect.ValueOf("Bar"))
		}
		fmt.Println(v.Field(0).Interface()) // Bar
		x := v.Type()
		for i := 0; i < x.NumField(); i++ {
			field := x.Field(i)
			fmt.Println(field.Tag)
		}
	}
}
