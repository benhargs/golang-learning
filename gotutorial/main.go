package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

//One line comment

/*
Multiline comment
Bonjour pal
*/

/*
Formatted print fmt.Printf()
The below are the different data types you need to call:

%s : string
%d : Integer
%c : Character
%f : FLoat
%o : Base 8
%x : Base 16
%v : Guesses based on data type
%T : Type of supplied value
*/

func main() {
	fmt.Printf("%s %d %c %f %t %o %x\n",
		"Stuff", 1, 'A', 3.14, true, 1, 1)
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.141592) //2 deciaml places, coming from the .2
	fmt.Printf("%9.f\n", 3.141592) //no deicmal places coming from the .

	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	pl("Hello", name)

	/*
			variables()
			cast()
			importantBirthdays()


		stringsEx()
		rune()
		mathsEx()
		forloop()
	*/
	//higherLower
}
