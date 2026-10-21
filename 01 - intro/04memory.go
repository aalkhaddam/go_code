/*

- variables are pointers in memory. They point to where the value is stored.
x = 5
x is a location with value 5
- reassigning x would change the value at location x
x = 2 
x is same location with value 2 now
- variable assignment with other variables creates a new location pointer, 
	but with the same value
y = x
y is a new location, but has the same value as x, 2 in this case.

why is this important?
- when you update a variable, if the variable itself is not updated, there 
	isn't anything that will point to that variable other than the variable 
	to change its value. 
example
x = 1
increment(x)

funx increment(x int){
	x++
}

This will still make x = 1. Not only is the increment not referencing the original variable,
	it's also not returning it so it basically does nothing here, right now. If you make it
	return a value, then put that value back into x with 'x =' or something, it will update
	the x value



*/