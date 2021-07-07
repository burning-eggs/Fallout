package main

import (
	"fmt"
	"os"
	"path/filepath"
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

func todosOfFile(path string) []Todo {
	// TODO: TodosOfFile not implemented
	return []Todo{
		Todo{
			Prefix:   "// ",
			Suffix:   "khooy",
			Id:       ref_str("#42"),
			Filename: "./main.go",
			Line:     10,
		},

		Todo{
			Prefix:   "// ",
			Suffix:   "foo",
			Id:       nil,
			Filename: "./src/foo.go",
			Line:     0,
		},
	}
}

func todosOfdir(dirpath string) ([]Todo, error) {
	result := []Todo{}

	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			for _, todo := range todosOfFile(path) {
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
