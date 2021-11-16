package cheapcash_test

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/aldy505/cheapcash"
)

func TestRead(t *testing.T) {
	rand := strconv.Itoa(rand.Int())
	c := cheapcash.Default()
	err := c.Write(rand, []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	v, err := c.Read(rand)
	if err != nil {
		t.Error("an error was thrown:", err)
	}
	if string(v) != "value" {
		t.Errorf("expected %s, got %v", "value", v)
	}
}

func TestRead_Conccurency(t *testing.T) {
	rand := strconv.Itoa(rand.Int())
	c := cheapcash.Default()

	err := c.Write(rand, []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	var wg sync.WaitGroup

	readFunc := func(){
		r, err := c.Read(rand)
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
