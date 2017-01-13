package main_test

import (
	"testing"

	"github.com/mukulrawat1986/GoActors"
	"github.com/stretchr/testify/assert"
)

func Test_FetchCredits_WithResults(t *testing.T) {
	a := assert.New(t)
	body := `{
		"cast": [
		{
			"id": 5966,
			"title": "Along Came Polly"
		},
		{
			"id": 1688,
			"name": "Friends"
		}
		]
		}`

	actor := &main.Actor{}
	FakeServer(body, func() {
		err := main.FetchCredits(actor)
		a.NoError(err)
		credits := actor.Credits
		a.Equal(2, len(credits))
		a.Equal("Along Came Polly", credits[0].NameOrTitle())
		a.Equal("Friends", credits[1].NameOrTitle())
	})
}

func Test_FetchCredits_WithOutResults(t *testing.T) {
	a := assert.New(t)
	body := `{
		"cast": []
		}`

	actor := &main.Actor{}
	FakeServer(body, func() {
		err := main.FetchCredits(actor)
		a.NoError(err)
		credits := actor.Credits
		a.Equal(0, len(credits))
	})
}

func Test_FilterCredits(t *testing.T) {
	a := assert.New(t)

	brad := main.Actor{Credits: []main.Credit{
		{ID: 1, Name: "Friends"},
		{ID: 2, Title: "World War Z"},
	}}

	jenn := main.Actor{Credits: []main.Credit{
		{ID: 3, Title: "Along Came Polly"},
		{ID: 1, Name: "Friends"},
		{ID: 4, Title: "The Break Up"},
	}}

	actors := []main.Actor{brad, jenn}

	credits := main.FilterCredits(actors)
	a.Equal(1, len(credits))
	a.Equal("Friends", credits[0].NameOrTitle())
}
