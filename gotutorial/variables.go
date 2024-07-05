package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
Don't need to call:
pl = fmt.Println
As it's already declared in hellogo.go,
which is also in package main.
*/

func variables() {
	pl := fmt.Println
	// var name type
	var vName string = "Benjy"
	var NumbyVar int = 49
	var DeciVar float32 = 10.1 //could be float64
	//float32 goes to +- 3.4e+38 (32 bits). float64 goes to +-1.7e+308 (64 bits)
	var multiExample1, multiExample2 = 16, 17
	var BoolVar = false
	RuneVar := "日本語"

	// Main types are: int, float64, bool, string, rune "Not sure what rune is"
	//To detect a type use the built in function below

	pl(reflect.TypeOf(vName))
	pl(reflect.TypeOf(NumbyVar))
	pl(reflect.TypeOf(BoolVar))
	pl(reflect.TypeOf(DeciVar))
	pl(reflect.TypeOf(RuneVar))      //rune is an int32 of characters and numbers, so you can incorporate symbols into your code.
	pl(multiExample1, multiExample2) // so the code doesn't fling up an error
}

func cast() {
	//This is a cast, it changes the variable type.
	castV1 := 1.5
	castV2 := int(castV1)
	pl(castV2) //from the output, cast rounds down on float to int.

	castV3 := -1.5
	castV4 := int(castV3)
	pl(castV4) //It rounds up for a negative number.

	cV5 := "500000000000"
	cV6, err := strconv.Atoi(cV5)
	pl(cV6, err, reflect.TypeOf(cV5))
	cV7 := 123
	cV8 := strconv.Itoa(cV7)
	pl(cV8, reflect.TypeOf(cV8))
}
