package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/fatih/color"
)

var inputReader = bufio.NewReader(os.Stdin)
var dictionary = []string{
	"Squirrel",
	"Statue of liberty",
	"Awkward",
	"Rhythm",
	"Microwave",
	"Galaxy",
	"Lucky",
	"Injury",
	"Vaporize",
	"Strenght",
	"Height",
	"Equip",
	"Keyhole",
	"Quartz",
	"Crystal",
	"Eiffel Tower",
	"Programming",
	"Restaurant",
	"Overload",
	"Hello World",
	"Painting",
	"Computer",
	"Window",
	"Laptop",
	"Mobile Phone",
	"Piggybank",
	"House",
	"Paintbrush",
	"Toothpaste",
	"Toothbrush",
	"Art Gallery",
	"Museum",
}

func main() {
	//Printing Instructions
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)
	boldRed.Println("Welcome to the game Hangman. The instructions are:")
	GameInstructions()
	fmt.Println()
	color.Black("The word is:")

	//Printing the random word
	rand.Seed(time.Now().UnixNano())
	chosenWord := getRandomWord()
	guessedLetters := initializeGuessedWords(chosenWord)
	hangmanState := 0

	for !isGameOver(chosenWord, guessedLetters, hangmanState) {
		printGameState(chosenWord, guessedLetters, hangmanState)
		input := readInput()
		if input == "exit" {
			fmt.Println("Goodbye...")
			os.Exit(0)
		} else if input == "Exit" {
			fmt.Println("Goodbye...")
			os.Exit(0)
		} else if len(input) != 1 {
			fmt.Println("Invalid input. Please use letters only.")
			continue
		}

		letter := rune(input[0])
		if isGuessCorrect(chosenWord, letter) {
			guessedLetters[letter] = true
		} else {
			fmt.Println("Wrong guess... Try again!")
			hangmanState++
		}
	}

	if isWordGuessed(chosenWord, guessedLetters) {

		t := color.New(color.FgGreen, color.Bold)
		t.Println("Game Over")
		fmt.Printf("The word was %s!\n", chosenWord)
		t.Println("You won!")

	} else if isHangmanComplete(hangmanState) {

		q := color.New(color.FgRed, color.Bold)
		fmt.Println(getHangmanDrawing(6))
		q.Println("Game Over")
		fmt.Printf("The word was %s...\n", chosenWord)
		q.Println("You lost!")

	} else {
		panic("invalid state. Game is over and there is no winner")
	}

}

func GameInstructions() {
	d := color.New(color.FgYellow)
	d.Printf(`
1. The player will select a letter from the alphabet. (Please use lowercase letters)
2. If the word/phrase contains that letter, all other letters equal to it are going to be revealed.
3. If the word/phrase doesnt contain this letter, a portion of the hangman is going to be added.
4. The game continues until:
a) the word/phrase is guessed and all letters are revealed - WINNER or,
b) all the parts of the hangman are displayed - LOSER
5. You can exit the program at all times by typing "Exit"/"exit" 
`)
}

func initializeGuessedWords(chosenWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(chosenWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(chosenWord[len(chosenWord)-1]))] = true

	return guessedLetters
}

func getRandomWord() string {
	chosenWord := dictionary[rand.Intn(len(dictionary))]
	return chosenWord
}

func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getHangmanDrawing(hangmanState))
}

func getWordGuessingProgress(chosenWord string, guessedLetters map[rune]bool) string {

	result := ""
	for _, ch := range chosenWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}

		result += " "
	}
	return result
}

func getHangmanDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d", hangmanState))
	if err != nil {
		panic(err)
	}

	return string(data)
}

func readInput() string {
	fmt.Print("> ")

	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func isGuessCorrect(chosenWord string, letter rune) bool {
	return strings.ContainsRune(chosenWord, letter)
}

func isGameOver(chosenWord string, guessedLetters map[rune]bool, hangmanState int) bool {
	return isWordGuessed(chosenWord, guessedLetters) || isHangmanComplete(hangmanState)
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}
	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 6
}
