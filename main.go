package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/Jihyun3478/logi-lang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Logi Programming language!\n", user.Username)
	fmt.Printf("Fell free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
