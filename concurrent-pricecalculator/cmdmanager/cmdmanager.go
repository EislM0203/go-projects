package cmdmanager

import "fmt"

type CmdManager struct {}

func (cm CmdManager) ReadLines() ([]string, error) {
	fmt.Println("Enter the input")

	var inputs []string

	for {
		var input string
		fmt.Print("input: ")
		fmt.Scanln(&input)
		if input == "0" {
			break
		}
		inputs = append(inputs, input)
	}

	return inputs, nil
}

func (cm CmdManager) WriteResult(job interface{}) error {
	fmt.Println(job)
	return nil
}

func New() *CmdManager {
	return &CmdManager{}
}