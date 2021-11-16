package cheapcash_test

import (
	"errors"
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

func TestAppend_NotExists(t *testing.T) {
	rand := strconv.Itoa(rand.Int())
	c := cheapcash.Default()

	err := c.Append(rand, []byte("Hello"))
	if err == nil {
		t.Error("expected an error, got nil")
	}

	if !errors.Is(err, cheapcash.ErrNotExists) {
		t.Errorf("expected %v, got %v", cheapcash.ErrNotExists, err)
	}
}
