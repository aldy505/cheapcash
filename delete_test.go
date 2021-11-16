package cheapcash_test

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/aldy505/cheapcash"
)

func TestDelete(t *testing.T) {
	randomValue := strconv.Itoa(rand.Int())
	c := cheapcash.Default()
	err := c.Write(randomValue, []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	err = c.Delete(randomValue)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}

func TestDelete_Concurrency(t *testing.T) {
	randomValue := strconv.Itoa(rand.Int())
	c := cheapcash.Default()
	err := c.Write(randomValue, []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	var wg sync.WaitGroup

	deleteFunc := func() {
		err = c.Delete(randomValue)
		if err != nil && !errors.Is(err, cheapcash.ErrNotExists) {
			t.Error("an error was thrown:", err)
		}
		wg.Done()
	}

	wg.Add(5)
	go deleteFunc()
	go deleteFunc()
	go deleteFunc()
	go deleteFunc()
	go deleteFunc()

	wg.Wait()
}
