package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
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

func create(flashcards []FlashCard, tracker strings.Builder) FlashCard {

	var flashcard FlashCard

	fmt.Println("The card:")
	tracker.WriteString("The card:")

	term := handleInput()
	tracker.WriteString(term)

	isDuplicated, _ := checkDuplicateTerm(term, flashcards)

	for isDuplicated {
		fmt.Printf("The card \"%s\" already exists. Try again: \n", term)
		tracker.WriteString(fmt.Sprintf("The card \"%s\" already exists. Try again: \n", term))
		term = handleInput()
		tracker.WriteString(term)
		isDuplicated, _ = checkDuplicateTerm(term, flashcards)
	}

	fmt.Println("The definition of the card:")
	tracker.WriteString("The definition of the card:")

	definition := handleInput()
	tracker.WriteString(definition)

	isDuplicated, _ = checkDuplicateDefinition(definition, flashcards)

	for isDuplicated {
		fmt.Printf("The definition \"%s\" already exists. Try again: \n", definition)
		tracker.WriteString(fmt.Sprintf("The definition \"%s\" already exists. Try again: \n", definition))

		definition = handleInput()
		tracker.WriteString(definition)

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

func ask(flashcards []FlashCard) bool {
	flashcardsLength := len(flashcards)

	if flashcardsLength == 0 {
		fmt.Println("No cards in-memory")
		return false
	}
	fmt.Println("How many times to ask?")
	times, err := strconv.Atoi(handleInput())
	if err != nil {
		fmt.Println("Not a valid number")
	}

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

	return true
}

func read(OGflashcards []FlashCard) []FlashCard {
	fmt.Println("File name:")
	fileName := handleInput()

	var flashcards []FlashCard

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("File not found.")
		return OGflashcards
	}

	err = json.Unmarshal(data, &flashcards)
	if err != nil {
		fmt.Println(err)
		return OGflashcards
	}

	for _, value := range flashcards {

		isDuplicated, key := checkDuplicateTerm(value.Term, OGflashcards)

		if isDuplicated {
			OGflashcards[key] = value
		} else {
			OGflashcards = append(OGflashcards, value)
		}
	}

	fmt.Printf("%d cards have been loaded. \n", len(flashcards))
	return OGflashcards

}

func export(flashcards []FlashCard) {
	fmt.Println("File name:")
	fileName := handleInput()

	data, err := json.Marshal(flashcards)
	if err != nil {
		log.Fatal(err) // exit the program if we have an unexpected error
	}

	if err = os.WriteFile(fileName, data, 0644); err != nil {
		log.Fatal(err) // exit the program if we have an unexpected error
	}

	fmt.Printf("%d cards have been saved. \n", len(flashcards))

}

func hardest(flashcards []FlashCard) bool {
	hardestCards := ""
	mistakesCount := 0
	multipleTerms := false
	for _, value := range flashcards {
		if value.Mistakes > mistakesCount {
			hardestCards = fmt.Sprintf("\"%s\"", value.Term)
			mistakesCount = value.Mistakes
		} else if value.Mistakes == mistakesCount && mistakesCount != 0 {
			multipleTerms = true
			hardestCards += fmt.Sprintf(", \"%s\"", value.Term)
		}
	}

	if mistakesCount == 0 {
		fmt.Println("There are no cards with errors.")
		return false
	}

	if multipleTerms {
		fmt.Printf("The hardest cards are %s", hardestCards)
	} else {
		fmt.Printf("The hardest card is %s. You have %d errors answering it", hardestCards, mistakesCount)
	}

	return true

}

func reset(flashcards []FlashCard) {
	for _, value := range flashcards {
		value.Mistakes = 0
	}

	fmt.Println("Card statistics have been reset.")
}

func tracking(tracker strings.Builder) {
	fmt.Println("File name: ")
	fileName := handleInput()

	if err := os.WriteFile(fileName, []byte(tracker.String()), 0644); err != nil {
		log.Fatal(err) // exit the program if we have an unexpected error
	}

	fmt.Println("The log has been saved.")
}

type FlashCard struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
	Mistakes   int    `json:"mistakes"`
}

func main() {

	action := ""
	var flashcards []FlashCard
	var tracker strings.Builder

	for action != "exit" {

		fmt.Println("Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats): ")
		action = handleInput()

		tracker.WriteString("Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats): ")
		tracker.WriteString(action)

		switch action {
		case "exit":
			break
		case "add":
			flashcard := create(flashcards, tracker)
			flashcards = append(flashcards, flashcard)

			fmt.Printf("The pair (\"%s\": \"%s\") has been added. \n", flashcard.Term, flashcard.Definition)
			tracker.WriteString(fmt.Sprintf("The pair (\"%s\": \"%s\") has been added. \n", flashcard.Term, flashcard.Definition))

			break
		case "remove":
			remove(&flashcards)
			break
		case "import":
			flashcards = read(flashcards)
			break
		case "export":
			export(flashcards)
			break
		case "ask":
			ask(flashcards)
			break
		case "log":
			tracking(tracker)
			break
		case "hardest card":
			hardest(flashcards)
			break
		case "reset stats":
			reset(flashcards)
			break

		default:
			fmt.Printf("`%s` is not a valid action. \n", action)
			tracker.WriteString(fmt.Sprintf("`%s` is not a valid action. \n", action))
		}

	}

	fmt.Println("Bye bye!")
	tracker.WriteString("Bye bye!")

}
