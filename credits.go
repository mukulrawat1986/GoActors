package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Credit struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

func (c Credit) NameOrTitle() string {
	if c.Name != "" {
		return c.Name
	}
	return c.Title
}

type CreditSearchResults struct {
	Cast []Credit `json:"cast"`
}

func FetchCredits(actor *Actor) error {
	url := fmt.Sprintf("%s/person/%d/combined_credits?api_key=%s", ApiRoot, actor.ID, ApiKey)
	results := CreditSearchResults{}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(&results)
	if err != nil {
		return err
	}

	actor.Credits = results.Cast
	return nil
}

func FilterCredits(actors []Actor) []Credit {
	credits := []Credit{}

	a := actors[0]
	al := len(actors)

	m := sync.Mutex{}
	var w sync.WaitGroup
	length := len(a.Credits)

	w.Add(length)

	// wait.Wait(len(a.Credits), func(i int){

	for i := 0; i < length; i++ {
		go func(w *sync.WaitGroup, i int, m *sync.Mutex) {
			c := a.Credits[i]
			count := 1

			for _, ab := range actors[1:] {
				for _, ac := range ab.Credits {
					if ac.ID == c.ID {
						count += 1
						break
					}
				}
			}

			if count == al {
				m.Lock()
				credits = append(credits, c)
				m.Unlock()
			}
			w.Done()
		}(&w, i, &m)
	}

	w.Wait()

	return credits
}
