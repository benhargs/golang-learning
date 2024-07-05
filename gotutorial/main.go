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

func main() {
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
	*/

	stringsEx()
	rune()
}
