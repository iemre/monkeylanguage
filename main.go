package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	fmt.Println("vim-go")
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey language!\n", user.Username)
	fmt.Printf("Feel free to tupe in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
