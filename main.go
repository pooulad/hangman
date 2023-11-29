package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

const (
	api_url = "https://api.dictionaryapi.dev/api/v2/entries/en"
)

type WordHint []struct {
	Meanings []struct {
		Definitions []struct {
			Definition string `json:"definition"`
		} `json:"definitions"`
	} `json:"meanings"`
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("hangman game: ")

	banner, err := os.ReadFile("./program/banner.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("reading file failed. please try again or submit an issue"))
	}

	fmt.Fprint(os.Stdout, string(banner)+"\n")

	gameWords := []string{"golang"}

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

	var wordHint WordHint
	resp, err := http.Get(fmt.Sprintf("%s/%s", api_url, randomWord))
	if err != nil {
		wordHint[0].Meanings[0].Definitions[0].Definition = ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		wordHint[0].Meanings[0].Definitions[0].Definition = ""
	}

	err = json.Unmarshal(body, &wordHint)
	if err != nil {
		wordHint[0].Meanings[0].Definitions[0].Definition = ""
	}

	lives := 5

	blanks := []string{}
	for range randomWord {
		blanks = append(blanks, "_")
	}

	correctAnswers := 0

loop:
	for {
		fmt.Printf("❤️ %d, Word letters : %s - word definition : %s\n", lives, strings.Join(blanks, " "), wordHint[0].Meanings[0].Definitions[0].Definition)

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
				correctAnswers += 1
				fmt.Printf("❤️ %d,✅correct answers %d, Word: %s - you won.\n", lives, correctAnswers, randomWord)
				randomWord = gameWords[rand.Intn(len(gameWords))]

				blanks = []string{}
				for range randomWord {
					blanks = append(blanks, "_")
				}

				if correctAnswers >= 3 {
					fmt.Println("🤩Congratulations, you have completed the game🤩")
					os.Exit(0)
				}

				goto loop
			}
		}
	}
}
