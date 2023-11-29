package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("hangman game: ")

	banner, err := os.ReadFile("./program/banner.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("reading file failed. please try again or submit an issue"))
	}

	fmt.Fprint(os.Stdout,string(banner) + "\n")

	gameWords := []string{}

	wordlist, err := os.Open("./wordlist.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("reading file failed. please try again or submit an issue"))
	}

	fileScanner := bufio.NewScanner(wordlist)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		gameWords = append(gameWords, fileScanner.Text())
	}

	wordlist.Close()

	randomWord := gameWords[rand.Intn(len(gameWords))]
	lives := 5

	blanks := []string{}
	for range randomWord {
		blanks = append(blanks, "_")
	}

loop:
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
				fmt.Printf("💔 0, Word was: %s - sorry your character is dead ☠️.\n", randomWord)
				break
			}

			if randomWord == strings.Join(blanks, "") {
				fmt.Printf("❤️ %d, Word: %s - you won.\n", lives, randomWord)
				randomWord = gameWords[rand.Intn(len(gameWords))]

				blanks = []string{}
				for range randomWord {
					blanks = append(blanks, "_")
				}

				goto loop
			}
		}
	}
}
