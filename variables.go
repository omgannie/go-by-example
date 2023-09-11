package main

import "fmt"

func main() {
	// Basic:
	var a = "initial"
	fmt.Println(a)
	
	// Delcare multiple variables at once.
	var b, c int = 1, 2
	fmt.Println(b, c)
	
	var d = true
	fmt.Println(d)
	
	// Variables w/o initial value is 0.
	var e int
	fmt.Println(e)
	
	// Reassignment:
	var f = "first"
	fmt.Println(f)
	f = "second"
	fmt.Println(f)
	
	// Reassignment to different type will fail:
	f = 20
	fmt.Println(f)
	
	// Shorthand. This is only available INSIDE of functions.
	g := "apple"
	fmt.Println(g)
}