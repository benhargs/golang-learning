package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func stringsEx() {
	strV1 := "A word"
	pl(strV1)
	replacer := strings.NewReplacer("A", "Another one... DJ Khalid") // Replaces the string on the left with the one on the right... If it exists
	strV2 := replacer.Replace(strV1)
	pl(strV2)

	//converts the length of the string into an int.
	pl("Length of the string: ", len(strV2))

	pl("Contains Another: ", strings.Contains(strV2, "Another")) //Returns a bool if the string on the right is contained in the string/variable on the left.

	pl("o index : ", strings.Index(strV2, "o")) //Finds the index of the first character that is being searched for (on the right of the comma), on the string on the left.

	pl("Replace: ", strings.Replace(strV2, "o", "0", -1))

	strV3 := "\nSome Words        .    \n" // \t let's you put a tab in your string \"" let's you put quotes in a string.
	pl(len(strV3))
	strV3 = strings.TrimSpace(strV3)
	pl(len(strV3), strV3)
	pl("Split: ", strings.Split("a-b-c-d", "-"))
	pl("Lower: ", strings.ToLower(strV2))
	pl("Upper: ", strings.ToUpper(strV2))
	pl("Prefix: ", strings.HasPrefix("tacocat", "taco")) //bool
	pl("Suffix: ", strings.HasSuffix("tacocat", "cat"))  //bool
}

func rune() {
	rStr := "abcdef"
	pl("Rune Count: ", utf8.RuneCountInString(rStr))

	for i, runeVal := range rStr {
		fmt.Printf("%d: %#U : %c\n", i, runeVal, runeVal) //formatted printing, just makes it easy to place different values within a string.
	}
}
