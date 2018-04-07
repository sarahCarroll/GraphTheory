# GraphTheory

## Problem statement

You must write a program in the Go programming language [2] that can
build a non-deterministic finite automaton (NFA) from a regular expression,
and can use the NFA to check if the regular expression matches any given
string of text. You must write the program from scratch and cannot use the
regexp package from the Go standard library nor any other external library.
A regular expression is a string containing a series of characters, some
of which may have a special meaning. For example, the three characters
“.”, “|”, and “∗
” have the special meanings “concatenate”, “or”, and “Kleene
star” respectively. So, 0.1 means a 0 followed by a 1, 0|1 means a 0 or a 1,
and 1∗ means any number of 1’s. These special characters must be used in
your submission.
Other special characters you might consider allowing as input are brackets
“()” which can be used for grouping, “+” which means “at least one of”, and
“?” which means “zero or one of”. You might also decide to remove the
concatenation character, so that 1.0 becomes 10, with the concatenation
implicit.
You may initially restrict the non-special characters your program works
with to 0 and 1, if you wish. However, you should at least attempt to expand
these to all of the digits, and the characters a to z, and A to Z.
You are expected to be able to break this project into a number of smaller
tasks that are easier to solve, and to plug these together after they have been
completed. You might do that for this project as follows:
1. Parse the regular expression from infix to postfix notation.
2. Build a series of small NFA’s for parts of the regular expression.
3. Use the smaller NFA’s to create the overall NFA.
4. Implement the matching algorithm using the NFA.
Overall your program might have the following layout.

## Breakdown
When creating this project I broke it down into the shunting yard algorithm and
Thompsons construction. The Shunting Yard Algorithm changed infix regular expressions 
to postfix regular expressions. Thompsons Construction determines what each character
represents and the action that occurs when it is reached.

## How to Run the program
In order to run the program you must Clone this github repository.

**git clone https://github.com/sarahCarroll/GraphTheory.git**

locate the folder it is in in the command prompt.Then you must run the command 

**go run Struct.go**

## References
https://en.wikipedia.org/wiki/Shunting-yard_algorithm
http://jacobappleton.io/2015/07/02/regex-ii-the-shunting-yard-algorithm/
https://gobyexample.com/range
http://www.cs.may.ie/staff/jpower/Courses/Previous/parsing/node5.html
https://golang.org/pkg/regexp/
https://swtch.com/~rsc/regexp/regexp1.html
https://web.microsoftstream.com/video/d08f6a02-23ec-4fa1-a781-585f1fd8c69e
https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
https://web.microsoftstream.com/video/946a7826-e536-4295-b050-857975162e6c
https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
