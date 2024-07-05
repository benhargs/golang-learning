package main

func importantBirthdays() {
	//Conditional Operators: >, <, >=, <=, ==, !=
	// Logical Operators: && (And), || (Or), ! (Not)

	iAge := 8

	if (iAge >= 1) && (iAge < 19) {
		pl("Important Birthday!")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday!")
	} else if iAge >= 65 {
		pl("Important Birthday!")
	} else {
		pl("Not an Important Birthday!")
	}
	pl("!true =", !true)

	qAge := 32
	if (qAge >= 1) && (qAge <= 18) || (qAge == 21) || (qAge == 50) || (qAge >= 65) {
		pl("Important Birthday!")
	} else {
		pl("Not an Important Birthday!")
	}
	pl("!true =", !true)
}
