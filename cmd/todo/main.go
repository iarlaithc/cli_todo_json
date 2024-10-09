package main

import (
	"flag"
	"fmt"
	todo "github/.com/iarlaithc/cli-todo-json"
	"os"
)

const (
	todoFile = ".todos.json"
)

func main() {
	// flags
	add := flag.Bool("add",false, "add a new todo")
	complete := flag.Int("complete", 0, "mark todo as complete")
	delete := flag.Int("delete", 0, "delete a todo")
	flag.Parse()

	// init todos 
	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	} 

	switch {
	case *add:
		todos.Add("Sample todo")
		err := todos.Store(todoFile)
		if err != nil{
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil{
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil{
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

	case *delete > 0:
		err := todos.Delete(*delete)
		if err != nil{
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.Store(todoFile)
		if err != nil{
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}