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

func Test_AskForNames(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Brad\nPitt\n")

	r := bytes.NewBuffer(b)

	main.AskForNames(r)

	a.Equal(len(main.ActorNames), 2)
	a.Equal(main.ActorNames[0], "Brad")
	a.Equal(main.ActorNames[1], "Pitt")
}

func Test_AskForNames_FourNames(t *testing.T) {
	setup()

	a := assert.New(t)
	b := []byte("Brad\nPitt\nJennifer\nLopez\n")

	r := bytes.NewBuffer(b)

	main.AskForNames(r)

	a.Equal(len(main.ActorNames), 4)
	a.Equal(main.ActorNames[0], "Brad")
	a.Equal(main.ActorNames[1], "Pitt")
	a.Equal(main.ActorNames[2], "Jennifer")
	a.Equal(main.ActorNames[3], "Lopez")
}
