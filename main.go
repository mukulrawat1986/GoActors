package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var ActorNames = []string{}

func Run(in stringReader, out io.Writer) {

	AskForNames(in)

	fmt.Fprintf(out, "You selected the following %d actors:\n", len(ActorNames))
	for _, v := range ActorNames {
		fmt.Fprintln(out, v)
	}
}

func main() {
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}
