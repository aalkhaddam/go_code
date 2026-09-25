package main

import "fmt"

func main(){
	messageFromDoris := []string{
		"You doing anything later?",
		"",
		"",
		"",
	}

	numMessages := float64(len(messagesfromDoris))
	costPerMessage := .02

	totalCost := costPerMessage * numMessages

	fmt.Println("Doris spent %.2f on text messages today \n", totalCost)
}

//what is compiling?
// you write code in readable text, computers only understand binary
// compiling is converting readable code to machine code/binary
// running "go build" turns code into executable

//distributing compiled program vs interpreted program
// interpreted - python. They just take your code and run it via "python main.py"
// downsides: they need python installed, and knowing how to use the tools
//   youre also giving them the code, and they can do what they want
// compiled - create exe, like a game, where the code is not visible
// they dont need to do anything else, they just run that program now

// deploying to server? generally easier to run compiled code to backend server
// because there are no runtime language dependecies 

// static typing - variables only have 1 type. Go enforces this
// pro: feedback on errors earlier
// trying to compile and run code, where an int is by itself in a fmt.Println 
// will throw an error. Needs to be a string, so the variable/data needs to change

// Go is a GarbageCollecting language, so memory management is autoamted, but it does not
//	have a virtual machine, like Java and JVM. 
// it has a runtime to manage GC - Java uses lots of memory, Rust uses less (since you specify memory), 
// Go is in the middle
// Depending on server load, Go memory would go up.

// uint - cant be negative
// number at the end of variable types tells you how many bits it's using. 
// can use less memory? use one with less bits

// byte == uint8
// rune == int32 - represents a unicode code point. Like a character in a string

var number int // initilization of number as int
var pi float64 = 3.14159

// shorthand
empty := "" // Go knows this is a string
// Still a static type
// depending on your machine, int will initialize 32 or 64 depending on computer's architecture (32bit, 64bit)
//recommended to only use these 4 unless you have a use case: int, uint, float64, complex128 (if youre using complex numbers)
// want to initialize with a different value? need to use longhand

//multiple on one line
number1, name1 := 10, "Alex"

//conversion
tempInt = 60
tempFloat = float64(tempInt)

//constants are not supported by shorthand

//String formatting
// 2 ways: Printf (prints a formatted string to std out), and Sprintf(returns the formatted string)
fmt.Printf("I am %v years old", 26)
message = fmt.Sprintf("He is %.2f minutes late", 38.90832423)
//%v is the value after. Could even be a string, value is ambiguous
//%s is a string, %d interpolates int as decimal (instead of binary for example), %f interpolates a decimal
fmt.Printf("I am %.2f years old", 26.3298749203874)
// goes to 2 decimal points