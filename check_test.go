package cheapcash_test

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/aldy505/cheapcash"
)

func TestExists(t *testing.T) {
	randomValue := strconv.Itoa(rand.Int())
	c := cheapcash.Default()
	b, err := c.Exists(c.Path + "/" + randomValue)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if b == true {
		t.Error("expected false, got true")
	}

	err = c.Write(randomValue, []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	b, err = c.Exists(c.Path + "/" + randomValue)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if b == false {
		t.Error("expected true, got false")
	}
}

func TestExists_Concurrency(t *testing.T) {
	randomValue := strconv.Itoa(rand.Int())
	c := cheapcash.Default()

	var wg sync.WaitGroup

	existsFunc := func() {
		b, err := c.Exists(c.Path + "/" + randomValue)
		if err != nil {
			t.Error("an error was thrown:", err)
		}
		if b == true {
			t.Error("expected false, got true")
		}
		wg.Done()
	}

	wg.Add(5)
	go existsFunc()
	go existsFunc()
	go existsFunc()
	go existsFunc()
	go existsFunc()

	wg.Wait()
}
