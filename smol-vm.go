package main

import (
	l "log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var log = l.New(os.Stdout, "", l.LUTC)

func main() {
	// getting file name
	programName := filepath.Base(os.Args[0])

	// checking for the filename
	if len(os.Args) < 2 {
		log.Printf("Usage %s <filename>", programName)
		return
	}

	filePath := os.Args[1]

	program, err := os.ReadFile(filePath)

	if err != nil {
		log.Printf("Error reading the file %s", err)
		return
	}

	instructions := strings.Split(string(program), "\n")

	for i := 0; i < len(instructions); i++ {
		instructions[i] = strings.TrimRight(instructions[i], "\r")
		instructions[i] = strings.TrimSpace(instructions[i])
	}

	// labels along with their line numbers
	labels := make(map[string]int)

	for lineNum, line := range instructions { // first pass to get the labels
		splittedLine := strings.Split(line, " ")
		op := splittedLine[0]
		data := strings.Join(splittedLine[1:], " ")

		if op == OP_LABEL {
			labels[data] = lineNum
		}
	}

	// variables names and their values
	variables := make(map[string]interface{})

	// stack to process the instructions
	stack := CreateNewStack[StackItem](len(instructions))

	// stack pointer to track which instruction is geing executed
	sp := -1

	for i := 0; i < len(instructions); i++ {
		line := instructions[i]

		splittedLine := strings.Split(line, " ")
		op := splittedLine[0]
		data := strings.Join(splittedLine[1:], " ")

		switch op {
		case OP_HALT:
			// halt the running program
			return
		case OP_SHOW:
			log.Print(data)
		case OP_PUSH:
			val, err := strconv.Atoi(data)
			if err != nil {
				log.Fatalf("Error converting string to number for %s with %s", data, err)
			}
			stack.Push(StackItem{
				name:  "constant",
				value: val,
			})
		case OP_POP:
			stack.Pop()
		case OP_PRINT:
			// pop the stack and print
			stackItem := stack.Peek()
			log.Printf("%d", GetValue(stackItem.value))
		case OP_LVALUE:
			// creating a variable and equating it to 0
			var memAdress = new(int)

			stack.Push(StackItem{
				name:  data,
				value: memAdress,
			})

			variables[data] = memAdress
		case OP_ASSIGN:
			// assign the varibles to the values
			first := stack.Pop()
			second := stack.Pop()
			// getting value
			newVal := GetValue(first.value)
			// getting memory address
			addr := second.value.(*int)
			*addr = newVal
			stack.Push(StackItem{
				name:  first.name,
				value: addr,
			})
		case OP_RVALUE:
			// get the value from the address of the variable and push it onto the stack
			address, exists := variables[data]
			tempData := 0

			if !exists {
				// create a variable named data and assign 0 to it
				variables[data] = tempData
				address = tempData
			}

			stack.Push(StackItem{
				name:  data,
				value: address,
			})

		case OP_GOTO:
			// go to the lable line and start executing from there
			sp = i
			labelLine, exists := labels[data]

			if !exists {
				log.Fatalf("No label named %s exists", data)
			}

			// moving the current index to the lable line to process from there
			i = labelLine
		case OP_BEGIN:
			// nothing as of now
		case OP_CALL:
			// call
		case OP_END:
			// end
		case OP_RETURN:
			i = sp
		case OP_ADD:
			first := stack.Pop()
			second := stack.Pop()

			stack.Push(StackItem{
				name:  "constant",
				value: GetValue(second.value) + GetValue(first.value),
			})
		case OP_DIVIDE:
			first := stack.Pop()
			second := stack.Pop()

			stack.Push(StackItem{
				name:  "constant",
				value: GetValue(second.value) / GetValue(first.value),
			})
		case OP_MODULO:
			first := stack.Pop()
			second := stack.Pop()

			stack.Push(StackItem{
				name:  "constant",
				value: GetValue(second.value) % GetValue(first.value),
			})
		case OP_SUBRACT:
			first := stack.Pop()
			second := stack.Pop()

			stack.Push(StackItem{
				name:  "constant",
				value: GetValue(second.value) - GetValue(first.value),
			})
		case OP_MULTIPLY:
			first := stack.Pop()
			second := stack.Pop()

			stack.Push(StackItem{
				name:  "constant",
				value: GetValue(second.value) * GetValue(first.value),
			})
		case OP_AND:
			first := stack.Pop()
			second := stack.Pop()

			stack.Push(StackItem{
				name:  "constant",
				value: GetValue(first.value) & GetValue(second.value),
			})
		case OP_OR:
			first := stack.Pop()
			second := stack.Pop()

			stack.Push(StackItem{
				name:  "constant",
				value: GetValue(first.value) | GetValue(second.value),
			})
		case OP_NOT:
			first := stack.Pop()

			firstVal := GetValue(first.value)
			result := 0

			if firstVal == 0 {
				result = 1
			} else {
				result = 0
			}

			stack.Push(StackItem{
				name:  "constant",
				value: result,
			})
		case OP_NOT_EQUAL:
			first := stack.Pop()
			second := stack.Pop()

			firstVal := GetValue(first.value)
			secondVal := GetValue(second.value)

			result := 0

			if firstVal != secondVal {
				result = 1
			} else {
				result = 0
			}

			stack.Push(StackItem{
				name:  "constant",
				value: result,
			})
		case OP_LESS_EQUAL:
			first := stack.Pop()
			second := stack.Pop()

			firstVal := GetValue(first.value)
			secondVal := GetValue(second.value)

			result := 0

			if firstVal <= secondVal {
				result = 0
			} else {
				result = 1
			}

			stack.Push(StackItem{
				name:  "constant",
				value: result,
			})

		case OP_GREAT_EQUAL:
			first := stack.Pop()
			second := stack.Pop()

			firstVal := GetValue(first.value)
			secondVal := GetValue(second.value)

			result := 0

			if firstVal >= secondVal {
				result = 0
			} else {
				result = 1
			}

			stack.Push(StackItem{
				name:  "constant",
				value: result,
			})
		case OP_LESS_THAN:
			first := stack.Pop()
			second := stack.Pop()

			firstVal := GetValue(first.value)
			secondVal := GetValue(second.value)

			result := 0

			if firstVal < secondVal {
				result = 0
			} else {
				result = 1
			}

			stack.Push(StackItem{
				name:  "constant",
				value: result,
			})
		case OP_GREATER_THAN:
			first := stack.Pop()
			second := stack.Pop()

			firstVal := GetValue(first.value)
			secondVal := GetValue(second.value)

			result := 0

			if firstVal > secondVal {
				result = 0
			} else {
				result = 1
			}

			stack.Push(StackItem{
				name:  "constant",
				value: result,
			})
		default:
			log.Printf("Invalid Operation %s", op)
		}

	}

}
