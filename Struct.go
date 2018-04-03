//Sarah Carroll- g00330821
//https://golang.org/pkg/regexp/
//https://swtch.com/~rsc/regexp/regexp1.html

package main

import (
	"fmt"
)

// the intPost function converts infix regular expressions to postfix regular expressions
// it breaks them down to the simplist form for the computer to understand
//using an algorithm called shunting yard- needing to use stacks in go
//https://en.wikipedia.org/wiki/Shunting-yard_algorithm
//http://jacobappleton.io/2015/07/02/regex-ii-the-shunting-yard-algorithm/
func intPost(infix string) string {
	//map to keep track +of special characters
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

type state struct {
	// store an indivual letter with the state / symbol of type rune
	// if not assigned with value gets 0 value(binary value 0)
	symbol rune

	// the max number of arrows coming from a state is 2
	// If an arrow other then an epsola is coming from any state thats the only arrow coming from that state
	// *=  pointer to the other state(linked list idea)
	edge1 *state
	edge2 *state
}

// build a list of state structs all collected together doing same thing as inputted collection
// In thompsons law you always have a single initial state and a single accept states

type nfa struct {
	// keeps track of the initial state and accept state of the fragement of the non-deterministic finite attomitons
	// There will always be one initial state but possibly numberous accept states.
	// speeds up search for initial and accept states (helper struct)
	initial *state
	accept  *state
}

// poregtonfa =post fix regular expression to non deterministic finite attomiton
// must return a pointer to nfa struct
//http://www.cs.may.ie/staff/jpower/Courses/Previous/parsing/node5.html
func poregtonfa(pofix string) *nfa {
	// thompsons algorithm works oon a stack ultimatily having a single fragment
	// on the stack which finnished which represents and nfa matching the regular expression
	nfaStack := []*nfa{}

	//search though the post fix regular expression one character at a time
	for _, r := range pofix {
		//for each special character we will either pop 1 or two fragments off our nfa stack
		switch r {
		//when r is concatinate character
		case '.':
			//pop 2 things off stack
			//index of last thing on stack
			//frag2 before frag 1 because the last thing on the stack/first thing to be taken off will be frag1
			frag2 := nfaStack[len(nfaStack)-1]
			//get everything on the nfa stack up to but not including the last item. : = upto but not including
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			// now frag1 and frag2 are two pointers to nfa fragments
			//want to join them together and push it back onto the nfa stack

			frag1.accept.edge1 = frag2.initial

			//set append to nfaStack a new pointer to an nfa struct that shows the new bigger nfa fragment
			//& needed because nfa is a struct - giving the address of the instance
			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})

		//when r is or
		case '|':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			//initial is a new state where edge1 points at frag1. initial and edge2 points at frag2.initial
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			//create 2 new states-accept adn initial and join those two states to the fragements that you pop off the stack
			accept := state{}

			//need to go in and get frag1.accepts edge which points at the accept state
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		//when r is the claney star
		case '*':
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}

			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			//old fragment with 2 extra states a new accept state and a new initial state
			//the new initial state points at the initial state of the fragment you popped off and the new accept state
			//the old accept state points at its own initial state and the new accept state
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		//default for when r isnt one of the 3 special characters
		default:
			//all you need to do is push the non special character to the stack , dont need to pull anything from the stack

			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		} //end switch
	} //end forloop

	//return the last/ only thing left on the stack which is the actually nfa you want.
	//***check that there is only one thing left on stack****
	return nfaStack[0]
}

//the function pomatch takes in 2 strings as its argument, the postfix regular expression and a string for it to be compared with.
//returning back a boolean-true/false
func pomatch(po string, s string) bool {

	//by default match is false/value you will return
	ismatch := false

	//Create a dfa from the postfix regular expression
	ponfa := poregtonfa(po)

	//when running an input string on a nfa-on nfa u can be in any number of states at a given time
	//create an array of pointers to state to keep track
	//dont want current to just have initial state but that and everything that can be got to
	//from the initial state -traverse the linked list that is poNfa
	current := []*state{}
	//keeps track of possible moves from your current state
	next := []*state{}

	//new function add state-helper function(recursive function in a non recursive form)
	//benifit - can be reeused to be called on next array
	//wat list of current states when we start off
	//in go if you want to pass an array and change it, you can convert into slice as shown below ([:])
	current = addState(current[:], ponfa.initial, ponfa.accept)

	//add to the next array the c state and ay state from the c state along e arrows
	//loop through string s a character at a time - current character is called r for rune
	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		//end of for loop(current)
		//every time you read a charater add all the possible next moves to next array
		//set current to next and empty the next array
		//when figured out the next set of states, make then the current set
		current, next = next, []*state{}
	}

	for _, c := range current {
		if c == ponfa.accept {
			//if state looping though on current array = accept state of nfa => match=true
			ismatch = true
			break
		} //end if c == ponfa.accept
	}

	//return ismatch
	return ismatch
}

//gets current array add state s and goes to s checking if its one the states with e arrows coming from it
func addState(l []*state, s *state, a *state) []*state {

	l = append(l, s)

	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
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

	nfa := poregtonfa("ab.c*|")
	// print out what is returned the nfa struct
	fmt.Println("postFix:      ", "a.(b.b)+.c")
	fmt.Println("nfa:         ", nfa)

	fmt.Println(pomatch(intPost("a.b.c*"), "abc"))

}
