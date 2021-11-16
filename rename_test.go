package cheapcash_test

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/aldy505/cheapcash"
)

func TestRename(t *testing.T) {
	randomValue := strconv.Itoa(rand.Int())
	randomValue2 := strconv.Itoa(rand.Int())

	c := cheapcash.Default()

	err := c.Write(randomValue, []byte("Another random value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	err = c.Rename(randomValue, randomValue2)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	v, err := c.Read(randomValue2)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if string(v) != "Another random value" {
		t.Errorf("expected %s, got %v", "Another random value", string(v))
	}

	_, err = c.Read(randomValue)
	if err != nil && !errors.Is(err, cheapcash.ErrNotExists) {
		t.Error("expected ErrNotExists, got:", err.Error())
	}
}

func TestRename_Concurrency(t *testing.T) {
	randomValue := strconv.Itoa(rand.Int())
	randomValue2 := strconv.Itoa(rand.Int())

	c := cheapcash.Default()
	err := c.Write(randomValue, []byte("Another random value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	var wg sync.WaitGroup

	renameFunc := func() {
		err = c.Rename(randomValue, randomValue2)
		if err != nil && !errors.Is(err, cheapcash.ErrNotExists) {
			t.Error("an error was thrown", err)
		}
		wg.Done()
	}

	wg.Add(5)
	go renameFunc()
	go renameFunc()
	go renameFunc()
	go renameFunc()
	go renameFunc()
	wg.Wait()
}

func TestRename_OldNotExists(t *testing.T) {
	randomValue := strconv.Itoa(rand.Int())

	c := cheapcash.Default()
	err := c.Rename(randomValue, "test")
	if !errors.Is(err, cheapcash.ErrNotExists) {
		t.Error("expected ErrNotExists, got:", err)
	}
}

func TestRename_NewExists(t *testing.T) {
	randomValue := strconv.Itoa(rand.Int())
	randomValue2 := strconv.Itoa(rand.Int())

	c := cheapcash.Default()
	err := c.Write(randomValue, []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	err = c.Write(randomValue2, []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	err = c.Rename(randomValue, randomValue2)
	if !errors.Is(err, cheapcash.ErrExists) {
		t.Error("expected ErrExists, got:", err)
	}
}
