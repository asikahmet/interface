package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Stringer interface {
	String() string
}
type S struct {
	V string
}

func (s *S) String() string {
	return s.V + "?"
}

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok {
		return v.String()
	}
	switch v := any.(type) {
	case int:
		return strconv.Itoa(v)
	}
	return ""
}

type Binary int

func (i Binary) String() string {
	return strconv.Itoa(i.Get())
}

func (i Binary) Get() int {
	reflect.TypeOf()
	return int(i)
}

func main() {
	fmt.Println(ToString(5))
	fmt.Println(ToString(&S{V: "V"}))
	fmt.Println(ToString(Binary(4)))
}
