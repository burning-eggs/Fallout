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
	// TODO(1): Todo.String doesn't print ID
	if todo.Id == nil {
		return fmt.Sprintf("%s:%d: %sTODO: %s\n", todo.Filename, todo.Line, todo.Prefix, todo.Suffix)
	} else {
		return fmt.Sprintf("%s:%d: %sTODO(%s): %s\n", todo.Filename, todo.Line, todo.Prefix, *todo.Id, todo.Suffix)
	}
}

func ref_str(x string) *string {
	return &x
}

func lineAsUnreportedTodo(line string) *Todo {
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

func lineAsReportedTodo(line string) *Todo {
	reportedTodo := regexp.MustCompile("^(.*)TODO\\((.*)\\): (.*)$")
	groups := reportedTodo.FindStringSubmatch(line)

	if groups != nil {
		return &Todo{
			Prefix:   groups[1],
			Suffix:   groups[3],
			Id:       &groups[2],
			Filename: "",
			Line:     0,
		}
	}

	return nil
}

func lineAsTodo(line string) *Todo {
	// TODO(7): lineAsTodo has false positive result inside of string literals
	if todo := lineAsUnreportedTodo(line); todo != nil {
		return todo
	}

	if todo := lineAsReportedTodo(line); todo != nil {
		return todo
	}

	return nil
}

func walkTodosOfFile(path string, visit func(Todo) error) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	line := 1
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		todo := lineAsTodo(scanner.Text())

		if todo != nil {
			todo.Filename = path
			todo.Line = line

			if err := visit(*todo); err != nil {
				return err
			}
		}

		line = line + 1
	}

	return scanner.Err()
}

func todosOfdir(dirpath string) ([]Todo, error) {
	result := []Todo{}

	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			err := walkTodosOfFile(path, func(todo Todo) error {
				result = append(result, todo)

				return nil
			})

			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}

func listSubCommand() {
	// TODO(4): listSubCommand doesn't handle error from todosOfDir
	todos, _ := todosOfdir(".")

	for _, todo := range todos {
		fmt.Printf("%v", todo)
	}
}

func reportSubCommand() {
	// TODO(5): reportSubCommand not implemented
	panic("Report is not implemented.")
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "list":
			listSubCommand()
		case "report":
			reportSubCommand()
		default:
			panic(fmt.Sprintf("`%s` Unknown Command", os.Args[1]))
		}
	} else {
		// TODO(6): Implement a map for options instead of printing them all
		fmt.Printf("fallout [option]\n\tlist: Lists all possible todos of a directory recursively\n\treport: Reports an issue to github\n")
	}
}
