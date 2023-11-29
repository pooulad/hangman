package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	gameWords := []string{"golang", "javascript", "php", "rust", "python"}
	randomWord := gameWords[rand.Intn(len(gameWords))]
	lives := 5

	blanks := []string{}
	for range randomWord {
		blanks = append(blanks, "_")
	}

	for {
		fmt.Printf("❤️ %d, Word letters : %s\n", lives, strings.Join(blanks, " "))

		var input string
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		if len(input) > len(randomWord) {
			fmt.Println("your input length is bigger than word length")
			lives--
		} else {
			for _, inputLetter := range input {
				correctAnswer := false

				for i, wordLetter := range randomWord {
					if inputLetter == wordLetter {
						blanks[i] = string(inputLetter)
						correctAnswer = true
					}
				}

				if !correctAnswer {
					lives--
				}
			}

			if lives <= 0 {
				fmt.Printf("💔 0, Word: %s - sorry you lost.\n", randomWord)
				break
			}

			if randomWord == strings.Join(blanks, "") {
				fmt.Printf("❤️ %d, Word: %s - you won.\n", lives, randomWord)
				break
			}
		}
	}
}
