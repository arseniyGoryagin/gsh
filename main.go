package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main(){

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to go shell")

	for {

		fmt.Print("> ")

		text, _ := reader.ReadString('\n')
		trimmedText := strings.Trim(text, "\n")
		splitText := strings.Split(trimmedText," ")
		command := splitText[0]
		args := splitText[1:]

		out, err := exec.Command(command, args...).Output()

		if(err != nil){
			fmt.Println(err)
		}

		stringOutput := string(out)

		fmt.Println(strings.Trim(stringOutput, "\n"))
	}

}