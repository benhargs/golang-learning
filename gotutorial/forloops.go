package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func forloop() {
	for x := 3; x >= 1; x-- {
		pl(x)
	}

	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}
}

func higherLower() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	fmt.Println("Guess between 1 and 50: ")
	number2Guess := rand.Intn(50) + 1
	counter := 0
	correctGuess := false

	for !correctGuess {
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guess = strings.TrimSpace(guess)
		turnGuess, err := strconv.Atoi(guess)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		counter++

		if turnGuess < number2Guess {
			fmt.Println("Guess Higher")
		} else if turnGuess > number2Guess {
			fmt.Println("Guess Lower")
		} else {
			correctGuess = true
			fmt.Printf("Congrats! You guessed %d in %d attempts.\n", number2Guess, counter)
		}
	}
}
