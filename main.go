package main

import "fmt"

func main() {
	// var name,age = "jahid",25 

	// fmt.Println("hello bangladesh")
	// fmt.Println("git add local to remote")
	// fmt.Println("my name is: ",name)
	// fmt.Println("my age is: ",age)
	//fmt.Printf("%d", 'ঔ')
	add(5,6)
	numToChar(65)
}

func add(x int, y int) {
	fmt.Println(x+y)

} 


func numToChar(x int){
	fmt.Printf("%c", x)
}