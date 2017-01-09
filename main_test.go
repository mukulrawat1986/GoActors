package main_test

import (
	"bytes"
	"testing"

	"github.com/mukulrawat1986/GoActors"

	"github.com/stretchr/testify/assert"
)

func Test_E2E(t *testing.T) {

	setup()

	a := assert.New(t)
	b := []byte("Brad Pitt\nJennifer Aniston\nn\n")

	r := bytes.NewBuffer(b)
	w := &bytes.Buffer{}

	main.Run(r, w)

	res := w.String()

	a.Contains(res, "You selected the following 2 actors")
	a.Contains(res, "Brad Pitt")
	a.Contains(res, "Jennifer Aniston")
}
