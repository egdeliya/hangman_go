package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

func initCountMap(word string) map[rune]int {
	lettersMap := map[rune]int{}
	for _, letter := range word {
		if _, ok := lettersMap[letter]; ok {
			lettersMap[letter] += 1
		} else {
			lettersMap[letter] = 1
		}
	}
	return lettersMap
}

func check(letter rune, lettersMap map[rune]int) (bool, bool) {
	acceptedLetter := false
	gameIsWon := false

	if count, ok := lettersMap[letter]; ok {
		acceptedLetter = true

		if count == 1 {
			delete(lettersMap, letter)
		} else {
			lettersMap[letter] -= 1
		}
		if len(lettersMap) < 1 {
			gameIsWon = true
		}
	}

	return acceptedLetter, gameIsWon
}

func currentState(curWordState string, isLetterInWord bool, letter uint8, word string) string {
	if isLetterInWord {
		for i := range word {
			if word[i] == letter && curWordState[i] != letter {
				if i < len(curWordState) {
					return curWordState[:i] + string(letter) + curWordState[i+1:]
				} else {
					return curWordState[:i] + string(letter)
				}
			}
		}
	}
	return curWordState
}

func hangman(stdin io.Reader, words []string) (gameIsWon bool) {
	maxTries := 5
	wastedTries := 0
	word := words[rand.Intn(len(words))]
	gameIsWon = false
	isInWord := false

	lettersMap := initCountMap(word)

	curWordState := strings.Repeat("*", len(word))
	reader := bufio.NewReader(stdin)

	for wastedTries < maxTries {
		print("Guess a letter: ")
		letter, _ := reader.ReadString('\n')
		if len(letter) > 1 && letter[len(letter)-1] != '\n' {
			wastedTries += 1
			continue
		}

		isInWord, gameIsWon = check(rune(letter[0]), lettersMap)

		curWordState = currentState(curWordState, isInWord, uint8(letter[0]), word)
		if isInWord {
			fmt.Println("Hit! ")
		} else {
			wastedTries += 1
			fmt.Printf("Missed, mistake %d out of %d. ", wastedTries, maxTries)
		}
		fmt.Printf("The word: %s\n\n", curWordState)

		if gameIsWon {
			return gameIsWon
		}

	}

	return gameIsWon
}

func main() {
	words := []string{"hello", "world"}
	gameIsWon := hangman(os.Stdin, words)

	if gameIsWon {
		fmt.Println("Win!")
	} else {
		fmt.Println("Lost :(")
	}
}
