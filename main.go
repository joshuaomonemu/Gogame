package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var playerScore, computerScore int

func main() {
	fmt.Println("Enter your name")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.Replace(name, "\n", "\n", -2)

	fmt.Printf("%s %s\n", "Welcome do you want to begin", name)

	for {
		fmt.Println(`Press "S" to start`)
		fmt.Println(`Press "E" to exit`)
		fmt.Println(`Press "V" to view score`)

		//reading single character from user

		teller := bufio.NewReader(os.Stdin)
		char, _, error := teller.ReadRune()

		if error != nil {
			fmt.Println("Error Occured ")
			fmt.Println(error)
			break
		}

		switch char {

		case 'S':
			fmt.Println("Game Starting...")
			startGame()
		case 'E':
			fmt.Println("Exiting...")
			os.Exit(1)
		case 'V':
			fmt.Println("%s  %d : %d  %s\n\n", name, playerScore, computerScore, "Computer")

		}
	}
}
func startGame() {
	fmt.Println(`Press "r" for rock 🥌`)
	fmt.Println(`Press "p" for paper 📃`)
	fmt.Println(`Press "s" for scissors ✂`)

	teller := bufio.NewReader(os.Stdin)
	userChoice, _, error := teller.ReadRune()

	if error != nil {
		fmt.Println("Error Occured ")
		fmt.Println(error)
		//break
	}

	computerChoice := getComputerChoice()
	choices := map[string]string{"r": "Rock", "p": "Paper", "s": "Scissors"}
	fmt.Println(("Computer choice was,"), choices[string(computerChoice)])
	fmt.Println(userChoice)
	results(userChoice, computerChoice)
}
func getComputerChoice() rune {
	choices := [3]rune{'r', 's', 'p'}
	return choices[rand.Intn(len(choices))]
}

func results(user rune, computer rune) {
	combinedresults := string(user) + string(computer)
	switch combinedresults {
	case "rs", "pr", "sp":
		fmt.Println("You Won 🎊🎊")
		playerScore++
	case "rp", "ps", "sr":
		fmt.Println("You lost 💔💔💔")
		computerScore++
	case "rr", "pp", "ss":
		fmt.Println("Its a tie👀👀")
	}
}
