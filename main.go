package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func handleInput() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func checkDuplicateDefinition(definition string, flashcards map[string]string) (bool, string) {
	for key, value := range flashcards {
		if value == definition {
			return true, key
		}
	}
	return false, ""
}

func checkDuplicateTerm(term string, flashcards map[string]string) bool {
	_, ok := flashcards[term]
	return ok
}

func add(flashcards map[string]string) (string, string) {

	fmt.Println("The card:")
	term := handleInput()
	isDuplicated := checkDuplicateTerm(term, flashcards)

	for isDuplicated {
		fmt.Printf("The card \"%s\" already exists. Try again: \n", term)
		term = handleInput()
		isDuplicated = checkDuplicateTerm(term, flashcards)
	}

	fmt.Println("The definition of the card:")
	definition := handleInput()
	isDuplicated, _ = checkDuplicateDefinition(definition, flashcards)

	for isDuplicated {
		fmt.Printf("The definition \"%s\" already exists. Try again: \n", definition)
		definition = handleInput()
		isDuplicated, _ = checkDuplicateDefinition(definition, flashcards)
	}

	return term, definition
}

func remove(flashcard *map[string]string) bool {
	var pFlashcard *map[string]string
	pFlashcard = flashcard

	fmt.Println("Which card?")
	card := handleInput()

	if isDuplicated := checkDuplicateTerm(card, *pFlashcard); !isDuplicated {
		fmt.Printf("Can't remove \"%s\": there is no such card. \n", card)
		return false
	}

	delete(*pFlashcard, card)
	fmt.Println("The card has been removed.")
	return true

}

func ask(flashcards map[string]string) {

	for key, value := range flashcards {
		fmt.Printf("Print the definition of \"%s\": \n", key)
		answer := handleInput()
		if answer == value {
			fmt.Println("Correct!")
		} else {
			duplicated, term := checkDuplicateDefinition(answer, flashcards)

			if duplicated {
				fmt.Printf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\". \n", value, term)
			} else {
				fmt.Printf("Wrong. The right answer is \"%s\". \n", value)
			}

		}
	}
}

func main() {

	action := ""
	flashcards := make(map[string]string)

	for action != "exit" {

		fmt.Println("Input the action (add, remove, import, export, ask, exit): ")
		action = handleInput()

		switch action {
		case "add":
			term, definition := add(flashcards)
			flashcards[term] = definition
			fmt.Printf("The pair (\"%s\": \"%s\") has been added. ", term, definition)
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
