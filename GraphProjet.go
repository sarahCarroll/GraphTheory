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
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	//poFix := ""
	poFix, s := []rune{}, []rune{}

	//creating algorithm
	//https://gobyexample.com/range - range , coverts sting to array of runes
	//_ is the index of the current character being read in, r represents the character
	for _, r := range infix {
		switch {
		case r == '(':
			s = append(s, r)

		case r == ')':
			//pop off the stack until we find (
			for s[len(s)-1] != '(' {
				poFix = append(poFix, s[len(s)-1])
				//get rid of last element on stack
				s = s[:len(s)-1]
			}
			s = s[:len(s)-1]

		//the current character is in the specials map
		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				poFix = append(poFix, s[len(s)-1])
				//get rid of last element on stack
				s = s[:len(s)-1]
			}
			s = append(s, r)
		default:
			//append adds the charater to the end of the poFix array
			poFix = append(poFix, r)
		}
	}

	for len(s) > 0 {
		//takes top element of stack and makes it top element of output(pofix)
		poFix = append(poFix, s[len(s)-1])
		//get rid of last element on stack
		s = s[:len(s)-1]
	}

	return string(poFix)
}

func main() {
	//following 4 examples of output are tesring the 3 basic operations | , * , .
	fmt.Println("Infix:      ", "a.b.c*")
	fmt.Println("postFix:    ", intPost("a.b.c*"))

	fmt.Println("Infix:      ", "(a.(b|d))*")
	fmt.Println("postFix:    ", intPost("(a.(b|d))*"))

	fmt.Println("Infix:      ", "a.(b|d).c*")
	fmt.Println("postFix:    ", intPost("a.(b|d).c*"))

	fmt.Println("Infix:      ", "a.(b.b)+.c")
	fmt.Println("postFix:    ", intPost("a.(b.b)+.c"))
}
