package cheapcash_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/aldy505/cheapcash"
)

func TestExists(t *testing.T) {
	rand := strconv.Itoa(rand.Int())
	c := cheapcash.Default()
	b, err := c.Exists(c.Path + "/" + rand)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if b == true {
		t.Error("expected false, got true")
	}

	err = c.Write(rand, []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	b, err = c.Exists(c.Path + "/" + rand)
	if err != nil {
		t.Error("an error was thrown:", err)
	}

	if b == false {
		t.Error("expected true, got false")
	}
}

func TestExists_Conccurency(t *testing.T) {
	rand := strconv.Itoa(rand.Int())
	c := cheapcash.Default()

	res := make(chan bool, 5)

	go func() {
		b, err := c.Exists(c.Path + "/" + rand)
		if err != nil {
			t.Error("an error was thrown:", err)
		}
		if b == true {
			t.Error("expected false, got true")
		}
		res <- true
	}()

	go func() {
		b, err := c.Exists(c.Path + "/" + rand)
		if err != nil {
			t.Error("an error was thrown:", err)
		}
		if b == true {
			t.Error("expected false, got true")
		}
		res <- true
	}()

	go func() {
		b, err := c.Exists(c.Path + "/" + rand)
		if err != nil {
			t.Error("an error was thrown:", err)
		}
		if b == true {
			t.Error("expected false, got true")
		}
		res <- true
	}()

	go func() {
		b, err := c.Exists(c.Path + "/" + rand)
		if err != nil {
			t.Error("an error was thrown:", err)
		}
		if b == true {
			t.Error("expected false, got true")
		}
		res <- true
	}()

	go func() {
		b, err := c.Exists(c.Path + "/" + rand)
		if err != nil {
			t.Error("an error was thrown:", err)
		}
		if b == true {
			t.Error("expected false, got true")
		}
		res <- true
	}()

	go func() {
		b, err := c.Exists(c.Path + "/" + rand)
		if err != nil {
			t.Error("an error was thrown:", err)
		}
		if b == true {
			t.Error("expected false, got true")
		}
		res <- true
	}()

	<-res
	<-res
	<-res
	<-res
	<-res
}
