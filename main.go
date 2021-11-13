package main

import (
	"bufio"
	"fmt"
	"os"
	cli "todo/cli"
	todo "todo/todo"
)

func main() {
	runProgram()
}

func runProgram() {
	id := 1
	ts := todo.TodoSlice{}
	reader := bufio.NewReader(os.Stdin)
	for {
		arg, in, err := cli.ReadConsole(reader)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch {
		case arg == "-h":
			cli.HelpCommandSlice()
		case arg == "-v":
			fmt.Println("0.1")
		case arg == "-a":
			err := todo.AddItem(&ts, &id, in)
			if err != nil {
				fmt.Println(err, "a")
			}
		case arg == "-l":
			ts.PrintItems()
		case arg == "-c":
			ts.PrintCompletedItems()
		case arg == "-m":
			tsp := &ts
			err := todo.MarkItemAsCompleted(tsp, in)
			if err != nil {
				fmt.Println(err)
			}
		case arg == "-d":
			tsp := &ts
			err := todo.DeleteItem(tsp, in)
			if err != nil {
				fmt.Println(err)

			}
		case arg == "-q":
			fmt.Println("Exiting the program...")
			os.Exit(1)
		}
	}
}
