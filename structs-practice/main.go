package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputable interface {
	saver
	Display()
}

// type outputable interface {
// 	Save() error
// 	Display()
// }

func main() {

	printSomething(1)
	printSomething(1.5)
	printSomething("hello")

	// title, content := getNoteData()
	// text := getTodoData()

	// note, err := note.New(title, content)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// todo, err := todo.New(text)
	// if err != nil {
	// 	fmt.Print(err)
	// }

	// note.Display()
	// err = saveData(note)
	// if err != nil {
	// 	return
	// }
	// todo.Display()
	// err = saveData(todo)
	// if err != nil {
	// 	return
	// }

	// outputData(note)
	// outputData(todo)
}

func printSomething(value interface{}) {

	typedVal, ok := value.(int)

	if ok {
		fmt.Println(typedVal + 1)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println(value)
	// }
}

func outputData(data outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}
	fmt.Println("Saving the note succeeded")
	return nil
}

func getTodoData() string {
	text := getUserInput("Enter your todo task : ")
	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note Title : ")
	content := getUserInput("Note Content : ")
	return title, content

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
