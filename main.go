package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleInput() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func checkDuplicateDefinition(definition string, flashcards []FlashCard) (bool, string) {
	for _, value := range flashcards {
		if value.Definition == definition {
			return true, value.Term
		}
	}
	return false, ""
}

func checkDuplicateTerm(term string, flashcards []FlashCard) (bool, int) {

	for key, value := range flashcards {
		if value.Term == term {
			return true, key
		}
	}
	return false, -1
}

func create(flashcards []FlashCard) FlashCard {

	var flashcard FlashCard

	fmt.Println("The card:")
	term := handleInput()
	isDuplicated, _ := checkDuplicateTerm(term, flashcards)

	for isDuplicated {
		fmt.Printf("The card \"%s\" already exists. Try again: \n", term)
		term = handleInput()
		isDuplicated, _ = checkDuplicateTerm(term, flashcards)
	}

	fmt.Println("The definition of the card:")
	definition := handleInput()
	isDuplicated, _ = checkDuplicateDefinition(definition, flashcards)

	for isDuplicated {
		fmt.Printf("The definition \"%s\" already exists. Try again: \n", definition)
		definition = handleInput()
		isDuplicated, _ = checkDuplicateDefinition(definition, flashcards)
	}

	flashcard.Term = term
	flashcard.Definition = definition

	return flashcard
}

func remove(flashcard *[]FlashCard) bool {
	var pFlashcard *[]FlashCard
	pFlashcard = flashcard

	fmt.Println("Which card?")
	card := handleInput()

	isDuplicated, key := checkDuplicateTerm(card, *pFlashcard)

	if !isDuplicated {
		fmt.Printf("Can't remove \"%s\": there is no such card. \n", card)
		return false
	}

	*pFlashcard = append((*pFlashcard)[:key], (*pFlashcard)[key+1:]...)

	fmt.Println("The card has been removed.")
	return true

}

func ask(flashcards []FlashCard) {

	fmt.Println("How many times to ask?")
	times, err := strconv.Atoi(handleInput())
	if err != nil {
		fmt.Println("Not a valid number")
	}

	flashcardsLength := len(flashcards)

	reseter := 0

	for i := 0; i < times; i++ {

		reseter++
		if reseter > flashcardsLength-1 {
			reseter = 0
		}

		fmt.Printf("Print the definition of \"%s\": \n", flashcards[reseter].Term)
		answer := handleInput()
		if answer == flashcards[reseter].Definition {
			fmt.Println("Correct!")
		} else {
			duplicated, term := checkDuplicateDefinition(answer, flashcards)

			if duplicated {
				fmt.Printf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\". \n", flashcards[reseter].Definition, term)
			} else {
				fmt.Printf("Wrong. The right answer is \"%s\". \n", flashcards[reseter].Definition)
			}

		}
	}
}

type FlashCard struct {
	Term, Definition string
}

func main() {

	action := ""
	var flashcards []FlashCard

	for action != "exit" {

		fmt.Println("Input the action (add, remove, import, export, ask, exit): ")
		action = handleInput()

		switch action {
		case "add":
			flashcard := create(flashcards)
			flashcards = append(flashcards, flashcard)
			fmt.Printf("The pair (\"%s\": \"%s\") has been added. \n", flashcard.Term, flashcard.Definition)
			break
		case "remove":
			remove(&flashcards)
			break
		case "import":
			fmt.Println("import")
			break
		case "export":
			fmt.Println("export")
			break
		case "ask":
			ask(flashcards)
			break
		case "exit":
			break

		default:
			fmt.Printf("`%s` is not a valid action. \n", action)
		}

	}

	fmt.Println("Bye bye!")

	//fmt.Print("Input the number of cards: \n")
	//quantity, _ = strconv.Atoi(handleInput())
	//
	//if quantity <= 0 {
	//	err := errors.New("not a valid quantity")
	//	fmt.Print(err)
	//}
	//
	//flashcards := buildingFlashCards(quantity)
	//
	//usingFlashCards(flashcards)

}
