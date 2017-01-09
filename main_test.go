package main_test

import (
	"bytes"
	"testing"

	"github.com/mukulrawat1986/GoActors"
	"github.com/stretchr/testify/assert"
)

func setup() {
	main.ActorNames = []string{}
}

func Test_AskForName(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Brad\n")

	r := bytes.NewBuffer(b)

	main.AskForName(r)

	a.Equal(len(main.ActorNames), 1)
	a.Equal(main.ActorNames[0], "Brad")
}
