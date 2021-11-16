package cheapcash_test

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/aldy505/cheapcash"
)

func TestRead(t *testing.T) {
	randomValue := strconv.Itoa(rand.Int())
	c := cheapcash.Default()
	err := c.Write(randomValue, []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	v, err := c.Read(randomValue)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	if string(v) != "value" {
		t.Errorf("expected %s, got %v", "value", v)
	}
}

func TestRead_Concurrency(t *testing.T) {
	randomValue := strconv.Itoa(rand.Int())
	c := cheapcash.Default()

	err := c.Write(randomValue, []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	var wg sync.WaitGroup

	readFunc := func() {
		r, err := c.Read(randomValue)
		if err != nil {
			t.Error("an error was thrown:", err)
		}
		if string(r) != "value" {
			t.Error("expected value, got:", string(r))
		}
		wg.Done()
	}

	wg.Add(5)
	go readFunc()
	go readFunc()
	go readFunc()
	go readFunc()
	go readFunc()

	wg.Wait()
}

func TestRead_NotExists(t *testing.T) {
	randomValue := strconv.Itoa(rand.Int())
	c := cheapcash.Default()

	_, err := c.Read(randomValue)
	if err == nil {
		t.Error("expected an error, got nil")
	}

	if !errors.Is(err, cheapcash.ErrNotExists) {
		t.Errorf("expected %v, got %v", cheapcash.ErrNotExists, err)
	}
}
