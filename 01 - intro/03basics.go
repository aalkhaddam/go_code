package main

import "fmt"

func main(){
	var1, var2, var3 := 1, 2, 3
	if var1 > 10{
		//some code
	} else if var2 > var3 {
		//some other code
	} else {
		fmt.Println("some text")
	}

	//shortened & scoped if statment:
	if length:=getLength(var); length>10{
		fmt.Println("length is enough")
	}
	// Benefits: 	the initialized variable in the "semi colon statment"
	//					 	is such that it is only within the if statement and 
	// 						is scoped to this statement. 
	// "safety hack" "limited scope"
}

//functions
func add(para1 int, para2 int) int {
	return para1 + para2
}
// return type is the only thing of note in the first line
// data type of parameters comes after the parameter name - same as regular vars

func subtr(para1, para2 int) int{
	return para1-para2
}
// if parameters are of the same data type, the type can be placed after the last one
// separated via comma
// adding other types? -> comma after data type and add
func age(para1, para2 int, name string) string{
	return fmt.Sprintf("Wow %v you are %v years old?", name, (para1+para2))
}

func f(func(int,int) int, int) int{
	return 6
}
// functions can be passed into others functions as parameters.
// NOTE: 	If the parameter function has 2 parameters and returns a type, then the 
// 				arguments that are passed into that (outer) function need to match
//				Can't pass a function that returns something different or has a 
//				different number of parameters

func a(x int) (int, int){
	return x-5, x+1
}
// multiple return values are wrapped with brackets
// NOTE: 	Go does not allow you to have empty variables.
// 				if you want to ignore a return value, you would use an underscore
x, _ := coordinates()
// coordinates() function returns 2 ints, but I only want the first value, the x
// (for example, I'm not going to use the y coordinate for a javelin contest. I don't
// care how high the spear went, just how far) so I use the underscore to let the 
// compiler know to omit that value from being read at this time. If I capture the y 
// value, but then I never use it ever again (so 'y' only shows up here), Go will 
// throw an error.

func b() (weight, volume, speed float64){
	return
}
// return values can also be named! 
// useful for when using multiple values or doing specific things
// naming them will initialize each value to their zero values
// leaving an empty return will automatically return the variables as they are at the return
// NOTE: this may hurt readability in long functions
// do your best to return the variables by name
