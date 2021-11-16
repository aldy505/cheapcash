package cheapcash_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/aldy505/cheapcash"
)

func TestAppend(t *testing.T) {
	rand := strconv.Itoa(rand.Int())
	c := cheapcash.Default()

	err := c.Write(rand, []byte("Hello"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	err = c.Append(rand, []byte("World"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	r, err := c.Read(rand)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if string(r) != "HelloWorld" {
		t.Errorf("expected %s, got %v", "HelloWorld", string(r))
	}
}
