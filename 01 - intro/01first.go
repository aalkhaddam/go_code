//https://www.youtube.com/watch?v=un6ZyFkqFKo

package main
// every file of go code has a package declaration
// package main: usually used because it will build into an executable go program
// "run it standalone"

import "fmt"
// importing fmt package 

func main(){ 
	fmt.Println("Hello World")
}
// every program starts execution at the top of the main function
// takes no arguments, doesn't return