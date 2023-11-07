package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type TuringMachine struct {
	tape       []rune
	head       int
	state      string
	transition map[TransitionKey]Transition
}

type TransitionKey struct {
	state     string
	readValue rune
}

type Transition struct {
	writeValue rune
	move       int
	newState   string
}

func NewTuringMachine(initialState string, input string) *TuringMachine {
	tape := []rune(input)
	return &TuringMachine{
		tape:       tape,
		head:       0,
		state:      initialState,
		transition: make(map[TransitionKey]Transition),
	}
}

func (tm *TuringMachine) AddTransition(state string, readValue rune, newState string, writeValue rune, move int) {
	key := TransitionKey{state, readValue}
	tm.transition[key] = Transition{writeValue, move, newState}
}

func (tm *TuringMachine) Run() {
	step := 0
	tapeString := string(tm.tape)
	tapeString = tapeString[:tm.head] + "[" + tm.state + "]" + tapeString[tm.head:]
	fmt.Printf("Step: %d, Tape: %s\n", step, tapeString)
	step++
	for {
		key := TransitionKey{tm.state, tm.tape[tm.head]}
		transition, exists := tm.transition[key]
		if !exists {
			break
		}

		tm.tape[tm.head] = transition.writeValue
		tm.head += transition.move
		tm.state = transition.newState

		if tm.head < 0 {
			tm.tape = append([]rune{'B'}, tm.tape...)
			tm.head = 0
		} else if tm.head >= len(tm.tape) {
			tm.tape = append(tm.tape, 'B')
		}

		tapeString := string(tm.tape)
		tapeString = tapeString[:tm.head] + "[" + tm.state + "]" + tapeString[tm.head:]
		fmt.Printf("Step: %d, Tape: %s\n", step, tapeString)
		step++
	}
}

func main() {

	var (
		Anfang = initialState(os.Args[1])
	)
	tm := NewTuringMachine(Anfang, os.Args[2])

	addTransitions(os.Args[1], tm)

	tm.Run()

}

func initialState(filename string) string {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	var Anfang = scanner.Text()

	return Anfang
}

func addTransitions(filename string, tm *TuringMachine) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()
	scanner.Scan()

	for scanner.Scan() {

		words := strings.Fields(scanner.Text())

		switch words[4] {
		case "R":
			words[4] = "1"
		case "L":
			words[4] = "-1"
		case "N":
			words[4] = "0"
		}

		int1, _ := strconv.ParseInt(words[4], 10, 6)

		state, readVal, nextState, writeVal, moveDir := words[0], []rune(words[1])[0], words[2], []rune(words[3])[0], int1
		tm.AddTransition(state, readVal, nextState, writeVal, int(moveDir))
	}
}
