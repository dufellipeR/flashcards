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

func handleInput(tracker *strings.Builder) string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	tracker.WriteString(strings.TrimSpace(line) + "\n")

	return strings.TrimSpace(line)
}

func trackedPrintln(a string, tracker *strings.Builder) {
	fmt.Println(a)
	tracker.WriteString(a + "\n")
}

func trackedPrintf(tracker *strings.Builder, format string, args ...interface{}) {
	fmt.Printf(format, args...)
	tracker.WriteString(fmt.Sprintf(format, args...))
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

func create(flashcards []FlashCard, tracker *strings.Builder) FlashCard {

	var flashcard FlashCard

	trackedPrintln("The card:", tracker)
	term := handleInput(tracker)

	isDuplicated, _ := checkDuplicateTerm(term, flashcards)

	for isDuplicated {
		trackedPrintf(tracker, "The card \"%s\" already exists. Try again: \n", term)
		term = handleInput(tracker)
		isDuplicated, _ = checkDuplicateTerm(term, flashcards)
	}

	trackedPrintln("The definition of the card:", tracker)
	definition := handleInput(tracker)

	isDuplicated, _ = checkDuplicateDefinition(definition, flashcards)

	for isDuplicated {
		trackedPrintf(tracker, "The definition \"%s\" already exists. Try again: \n", definition)
		definition = handleInput(tracker)

		isDuplicated, _ = checkDuplicateDefinition(definition, flashcards)
	}

	flashcard.Term = term
	flashcard.Definition = definition

	return flashcard
}

func remove(flashcard *[]FlashCard, tracker *strings.Builder) bool {
	var pFlashcard *[]FlashCard
	pFlashcard = flashcard

	trackedPrintln("Which card?", tracker)

	card := handleInput(tracker)

	isDuplicated, key := checkDuplicateTerm(card, *pFlashcard)

	if !isDuplicated {
		trackedPrintf(tracker, "Can't remove \"%s\": there is no such card. \n", card)
		return false
	}

	*pFlashcard = append((*pFlashcard)[:key], (*pFlashcard)[key+1:]...)

	trackedPrintln("The card has been removed.", tracker)
	return true

}

func ask(flashcards []FlashCard, tracker *strings.Builder) bool {
	flashcardsLength := len(flashcards)

	if flashcardsLength == 0 {
		trackedPrintln("No cards in-memory", tracker)
		return false
	}
	trackedPrintln("How many times to ask?", tracker)
	times, err := strconv.Atoi(handleInput(tracker))
	if err != nil {
		trackedPrintln("Not a valid number", tracker)
	}

	reseter := 0

	for i := 0; i < times; i++ {

		reseter++
		if reseter > flashcardsLength-1 {
			reseter = 0
		}

		trackedPrintf(tracker, "Print the definition of \"%s\": \n", flashcards[reseter].Term)
		answer := handleInput(tracker)
		if answer == flashcards[reseter].Definition {
			trackedPrintln("Correct!", tracker)
		} else {
			flashcards[reseter].Mistakes += 1
			duplicated, term := checkDuplicateDefinition(answer, flashcards)

			if duplicated {
				trackedPrintf(tracker, "Wrong. The right answer is \"%s\", but your definition is correct for \"%s\". \n", flashcards[reseter].Definition, term)
			} else {
				trackedPrintf(tracker, "Wrong. The right answer is \"%s\". \n", flashcards[reseter].Definition)
			}

		}
	}

	return true
}

func read(OGflashcards []FlashCard, tracker *strings.Builder) []FlashCard {
	trackedPrintln("File name:", tracker)
	fileName := handleInput(tracker)

	var flashcards []FlashCard

	data, err := os.ReadFile(fileName)
	if err != nil {
		trackedPrintln("File not found.", tracker)
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

func export(flashcards []FlashCard, tracker *strings.Builder) {
	trackedPrintln("File name:", tracker)
	fileName := handleInput(tracker)

	data, err := json.Marshal(flashcards)
	if err != nil {
		log.Fatal(err) // exit the program if we have an unexpected error
	}

	if err = os.WriteFile(fileName, data, 0644); err != nil {
		log.Fatal(err) // exit the program if we have an unexpected error
	}

	trackedPrintf(tracker, "%d cards have been saved. \n", len(flashcards))

}

func hardest(flashcards []FlashCard, tracker *strings.Builder) bool {
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
		trackedPrintln("There are no cards with errors.", tracker)
		return false
	}

	if multipleTerms {
		trackedPrintf(tracker, "The hardest card is %s. You have %d errors answering it \n", hardestCards, mistakesCount)
	} else {
		trackedPrintf(tracker, "The hardest card is %s. You have %d errors answering it \n", hardestCards, mistakesCount)

	}

	return true

}

func reset(flashcards *[]FlashCard, tracker *strings.Builder) {
	for key, _ := range *flashcards {
		(*flashcards)[key].Mistakes = 0
	}

	trackedPrintln("Card statistics have been reset.", tracker)
}

func tracking(tracker *strings.Builder) {
	trackedPrintln("File name: ", tracker)
	fileName := handleInput(tracker)

	if err := os.WriteFile(fileName, []byte(tracker.String()), 0644); err != nil {
		log.Fatal(err) // exit the program if we have an unexpected error
	}

	trackedPrintln("The log has been saved.", tracker)
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

		trackedPrintln("Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats): ", &tracker)
		action = handleInput(&tracker)

		switch action {
		case "exit":
			break
		case "add":

			flashcard := create(flashcards, &tracker)
			flashcards = append(flashcards, flashcard)

			trackedPrintf(&tracker, "The pair (\"%s\": \"%s\") has been added. \n", flashcard.Term, flashcard.Definition)

			break
		case "remove":
			remove(&flashcards, &tracker)
			break
		case "import":
			flashcards = read(flashcards, &tracker)
			break
		case "export":
			export(flashcards, &tracker)
			break
		case "ask":
			ask(flashcards, &tracker)
			break
		case "log":
			tracking(&tracker)
			break
		case "hardest card":
			hardest(flashcards, &tracker)
			break
		case "reset stats":
			reset(&flashcards, &tracker)
			break

		default:
			trackedPrintf(&tracker, "`%s` is not a valid action. \n", action)
		}
	}

	trackedPrintln("Bye bye!", &tracker)

}
