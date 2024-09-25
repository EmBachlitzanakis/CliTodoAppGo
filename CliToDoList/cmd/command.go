package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add      string
	Edit     string
	Priority string
	Del      int
	Toggle   int
	List     bool
	AIHelp   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title : Add the priority too")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and specify a new title. id:new_title")
	flag.StringVar(&cf.Priority, "prior", "", "Edit Priority a todo by index and specify a new Priority. id:priority")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to Delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.BoolVar(&cf.AIHelp, "AIHelp", false, "List all todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.AIHelp:
		todos.Aiprint()
	case cf.Add != "":
		parts := strings.SplitN(cf.Add, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for add, Please use  { new title:priority } ")
			os.Exit(1)
		}
		todos.add(parts[0], parts[1])

	case cf.Priority != "":
		parts := strings.SplitN(cf.Priority, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit, Please use id:priority")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit")
			os.Exit(1)
		}

		todos.editPriority(index, parts[1])

	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit, Please use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)
	default:
		fmt.Println("Invalid Command")
	}
}
