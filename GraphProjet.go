package main

import (
	"fmt"
)

// the intPost function converts infix regular expressions to postfix regular expressions
// it breaks them down to the simplist form for the computer to understand
//using an algorithm called shunting yard-needing to use stacks in go
//https://en.wikipedia.org/wiki/Shunting-yard_algorithm
func intPost(infix string) string {
    //map to keep track of special characters
    //if character is in specials do soemthing with it
    //map to keep track of presedence
    specials := map[rune]int{'*':10 ,'.':9 , '|':8}
	//poFix := ""
	poFix,s := []rune{}.[]rune{}
	

	return string(poFix)
}

func main() {
	//following 4 examples of output are tesring the 3 basic operations | , * , .
	fmt.Println("Infix:      ", "a.b.c*")
	fmt.Println("postFix:    ", intPost("a.b.c*"))

	fmt.Println("Infix:      ", "(a.b|d))*")
	fmt.Println("postFix:    ", intPost("(a.b|d))*"))

	fmt.Println("Infix:      ", "a.(b|d).c*")
	fmt.Println("postFix:    ", intPost("a.(b|d).c*"))

	fmt.Println("Infix:      ", "a.(b.b)+.c")
	fmt.Println("postFix:    ", intPost("a.(b.b)+.c"))
}
