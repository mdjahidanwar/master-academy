package main

import (
	"fmt"
	"reflect"
)

func main() {
	x,y,z:="jahid",20,1995
	fmt.Printf("%v=%T\n", x,x)
	fmt.Printf("%v=%T\n", y,y)
	fmt.Printf("%v=%T\n", z,z)
	fmt.Println(reflect.TypeOf(x))
}
