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

func AskForNames(r stringReader) {
	asking := true
	for asking {
		if len(ActorNames) < 2 {
			AskForName(r)
		} else {
			asking = false
		}
	}
}

func main() {
	AskForNames(bufio.NewReader(os.Stdin))

	fmt.Printf("You selected the following %d actors:\n", len(ActorNames))
	for _, v := range ActorNames {
		fmt.Println(v)
	}
}
