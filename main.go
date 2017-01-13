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

func init() {
	ApiKey = os.Getenv("TMDB_KEY")
}

func Run(in stringReader, out io.Writer) {

	ActorNames = []string{}
	AskForNames(in)

	actors := []Actor{}
	m := sync.Mutex{}

	fmt.Fprintf(out, "\nYou selected the following %d actors:\n", len(ActorNames))

	var w sync.WaitGroup
	length := len(ActorNames)
	w.Add(length)

	for i := 0; i < length; i++ {
		go func(w *sync.WaitGroup, i int, out io.Writer, m *sync.Mutex) {
			actor, err := FetchActor(ActorNames[i])
			if err != nil {
				log.Fatal(err)
			}
			m.Lock()
			actors = append(actors, actor)
			m.Unlock()
			fmt.Fprintln(out, actor.Name)
			w.Done()
		}(&w, i, out, &m)
	}

	w.Wait()

	credits := FilterCredits(actors)
	if len(credits) > 0 {
		fmt.Fprintln(out, "\nThey have appeared in the following movies and TV shows together:")
		for _, c := range credits {
			fmt.Fprintln(out, c.NameOrTitle())
		}
	} else {
		fmt.Fprintln(out, "\nHave not appeared in anything together.")
	}
}

func main() {
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}
