package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

type Todo struct {
	Prefix   string
	Suffix   string
	Id       *string
	Filename string
	Line     int
}

func (todo Todo) String() string {
	// TODO: Todo.String doesn't print ID
	if todo.Id == nil {
		return fmt.Sprintf("%s:%d: %sTODO: %s\n", todo.Filename, todo.Line, todo.Prefix, todo.Suffix)
	} else {
		return fmt.Sprintf("%s:%d: %sTODO(%s): %s\n", todo.Filename, todo.Line, todo.Prefix, *todo.Id, todo.Suffix)
	}
}

func ref_str(x string) *string {
	return &x
}

func lineAsTodo(line string) *Todo {
	// TODO: lineAsTodo doesn't support reported TODOs
	// TODO: lineAsTodo has false positive result inside of string literals
	unreportedTodo := regexp.MustCompile("^(.*)TODO: (.*)$")
	groups := unreportedTodo.FindStringSubmatch(line)

	if groups != nil {
		return &Todo{
			Prefix:   groups[1],
			Suffix:   groups[2],
			Id:       nil,
			Filename: "",
			Line:     0,
		}
	}

	return nil
}

func todosOfFile(path string) ([]Todo, error) {
	result := []Todo{}
	file, err := os.Open(path)

	if err != nil {
		return []Todo{}, err
	}

	defer file.Close()

	line := 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		todo := lineAsTodo(scanner.Text())

		if todo != nil {
			todo.Filename = path
			todo.Line = line

			result = append(result, *todo)
		}

		line = line + 1
	}

	return result, scanner.Err()
}

func todosOfdir(dirpath string) ([]Todo, error) {
	result := []Todo{}

	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			todos, err := todosOfFile(path)

			if err != nil {
				return err
			}

			for _, todo := range todos {
				result = append(result, todo)
			}
		}

		return nil
	})

	return result, err
}

func listSubCommand() {
	// TODO: listSubCommand doesn't handle error from todosOfDir
	todos, _ := todosOfdir(".")

	for _, todo := range todos {
		fmt.Printf("%v", todo)
	}
}

func reportSubCommand() {
	// TODO: reportSubCommand not implemented
	panic("Report is not implemented.")
}

func main() {
	// TODO: Index out of range error when no subcommands are provided
	switch os.Args[1] {
	case "list":
		listSubCommand()
	case "report":
		reportSubCommand()
	default:
		panic(fmt.Sprintf("`%s` Unknown Command", os.Args[1]))
	}
}
