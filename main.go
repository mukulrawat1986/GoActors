package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var ActorNames = []string{}

type stringReader interface {
	ReadString(byte) (string, error)
}

func AskForName(r stringReader) {
	fmt.Println("Please enter an actor's name: ")
	name, _ := r.ReadString('\n')
	name = strings.TrimSpace(name)
	ActorNames = append(ActorNames, name)
}

func main() {
	AskForName(bufio.NewReader(os.Stdin))
}
