package features

import (
	"fmt"
	"github.com/fatih/color"
)

var LOGO string = fmt.Sprintln(`┌────────────────────────────────────────────────────────────────────┐ 
•                       GoCron v1.0.4                                •
• █████████  █████████  ████████  █████████  █████████  ███      ██  •
• █          █       █  █         █       █  █       █  ████     ██  •
• █          █       █  █         █       █  █       █  ██ ██    ██  •
• █     ███  █       █  █         █████████  █       █  ██  ██   ██  •
• █       █  █       █  █         █    ██    █       █  ██   ██  ██  •
• █       █  █       █  █         █     ██   █       █  ██    ██ ██  •
• █████████  █████████  ████████  █      ██  █████████  ██     ████  •
•       is a API for learning GO language with example               •
└────────────────────────────────────────────────────────────────────┘ `)

var Message string = fmt.Sprintln(`
my name is Sina LalehBakhsh, I hope this API is useful for you
after running program, write your single word about any of GO language.
if your perpuse is more than one word, for convenience searching, just write keywords.
like this:
	map slice
	`)

func HelpMessage() {
	color.HiBlue(`
	┌─────────────────┐
	|  GoCron v1.0.4  |
	└─────────────────┘
	`)
	words := SplitIntoWords(Message)
	PrintWordByWord(words)
	fmt.Println()
	color.HiBlue(fmt.Sprintln(`You can use this command before start:
	◉============ Run
	1-
	go run .
	
	◉============ Build
	1-
	go build .
	2-
	./Gocron
	
	◉============ Help
	./Gocron -h 
	./Gocron help
	./Gocron -help
	./Gocron --help

	◉============ Show All
	./Gocron all
	./Gocron -all
	./Gocron --all
	`))
	
}

var TitleOfHelp = []string{
	"-h",
	"help",
	"-help",
	"--help",
}
