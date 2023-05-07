

# Flashcards GO

This is a command-line tool written in GO that allows you to create, manage, and use flashcards for studying.

## Features

- Add and remove flashcards.
- Use flashcards to study and keep track of incorrect answers.
- See the flashcard with the most incorrect answers.
- Reset flashcards statistics
- Import and export flashcards.
- Export a log of inputs

## Installation

To use this tool, you need to have GO installed on your machine. You can download it from the official website: https://golang.org/dl/.

Once you have installed GO, clone this repository and navigate to the project directory in your terminal.

```bash
git clone https://github.com/dufellipeR/flashcards.git
cd flashcards
```

Build the project using the following command:

```bash
go build main.go
```

This will create an executable file named `main` in the project directory.

```bash
./main
```

## Usage

To see the available commands, run:

```bash
./main --help
```

The tool has the following running commands:

- `add`: add a new flashcard to the deck.
- `remove`: remove a flashcard from the deck.
- `ask`: use the flashcards to study.
- `hardest card`: see the statistics for the deck.
- `reset stats`: reset the statistics for the deck.
- `import`: import flashcards from a file.
- `export`: export flashcards to a file.