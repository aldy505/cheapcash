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
	rand := strconv.Itoa(rand.Int())
	c := cheapcash.Default()
	err := c.Write(rand, []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	err = c.Delete(rand)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}

func TestDelete_Conccurency(t *testing.T) {
	rand := strconv.Itoa(rand.Int())
	c := cheapcash.Default()
	err := c.Write(rand, []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	var wg sync.WaitGroup

	deleteFunc := func() {
		err = c.Delete(rand)
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
