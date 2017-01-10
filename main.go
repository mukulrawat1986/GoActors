package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

var ActorNames = []string{}

func Run(in stringReader, out io.Writer) {

	AskForNames(in)

	fmt.Fprintf(out, "You selected the following %d actors:\n", len(ActorNames))

	var w sync.WaitGroup
	length := len(ActorNames)
	w.Add(length)
	for i := 0; i < length; i++ {
		go func(w *sync.WaitGroup, i int, out io.Writer) {
			actor, err := FetchActor(ActorNames[i])
			if err != nil {
				log.Panic(err)
			}
			fmt.Fprintln(out, actor.Name)
			w.Done()
		}(&w, i, out)
	}
	w.Wait()
}

func main() {
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}
