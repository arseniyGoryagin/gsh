package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const REDIRECT_ARG = ">f"

func isRedirectOutPutArg(arg string) bool {
	return arg == REDIRECT_ARG
}

func parseInputText(text string) (string, []string) {
	trimmedText := strings.Trim(text, "\n")
	splitText := strings.Split(trimmedText, " ")
	command := splitText[0]
	args := splitText[1:]
	return command, args
}

func parseExecOutput(output []byte) string {
	stringOutput := string(output)
	return strings.Trim(stringOutput, "\n")
}

func openFile(filename string) (*os.File, error){
	return os.Create(filename)
}

func writeToFile(file *os.File, data string){
	file.WriteString(data)
	file.Sync()
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to go shell")

	for {

		fmt.Print("> ")

		textInput, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input")
			continue
		}

		command, args := parseInputText(textInput)

		var redirecttoFilename string

		argsEnd := len(args)

		for i, arg := range args {

			if isRedirectOutPutArg(arg) {
				redirecttoFilename = args[i+1]
				argsEnd = i
				break;
			}

		}

		out, err := exec.Command(command, args[0:argsEnd]...).Output()

		if err != nil {
			fmt.Println(err)
			continue
		}

		textOutPut := parseExecOutput(out)

		if redirecttoFilename  == "" {
			fmt.Println(textOutPut)
		}

		file, err := openFile(redirecttoFilename)

		if err != nil {
			fmt.Println("Error writing to file")
			continue
		}

		writeToFile(file, textOutPut)

		file.Close()

	}

}
