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

func buildingFlashCards(quantity int) map[string]string {
	flashcards := make(map[string]string)

	for i := 0; i < quantity; i++ {

		fmt.Printf("The term for card #%d: \n", i+1)
		term := handleInput()
		isDuplicated := checkDuplicateTerm(term, flashcards)

		for isDuplicated {
			fmt.Printf("The term \"%s\" already exists. Try again: \n", term)
			term = handleInput()
			isDuplicated = checkDuplicateTerm(term, flashcards)
		}

		fmt.Printf("The definition for card #%d: \n", i+1)
		definition := handleInput()
		isDuplicated, _ = checkDuplicateDefinition(definition, flashcards)

		for isDuplicated {
			fmt.Printf("The definition \"%s\" already exists. Try again: \n", definition)
			definition = handleInput()
			isDuplicated, _ = checkDuplicateDefinition(definition, flashcards)
		}

		flashcards[term] = definition
	}

	return flashcards
}

func usingFlashCards(flashcards map[string]string) {

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

type FlashCard struct {
	Term, Definition string
}

func main() {

	action := ""

	for action != "exit" {

		fmt.Println("Input the action (add, remove, import, export, ask, exit): ")
		action = handleInput()

		switch action {
		case "add":
			fmt.Println("add")
			break
		case "remove":
			fmt.Println("remove")
			break
		case "import":
			fmt.Println("import")
			break
		case "export":
			fmt.Println("export")
			break
		case "ask":
			fmt.Println("ask")
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
