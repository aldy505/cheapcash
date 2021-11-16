package cheapcash_test

import (
	"sync"
	"testing"

	"github.com/aldy505/cheapcash"
)

func TestWrite(t *testing.T) {
	c := cheapcash.Default()
	err := c.Write("key", []byte("value"))
	if err != nil {
		t.Error("an error was thrown:", err)
	}
}

func TestWrite_Conccurency(t *testing.T) {
	c := cheapcash.Default()

  var wg sync.WaitGroup

  writeFunc := func() {
		err := c.Write("key1", []byte("value1"))
		if err != nil {
			t.Error("an error was thrown:", err)
		}
    wg.Done()
	}

	wg.Add(5)
  go writeFunc()
  go writeFunc()
  go writeFunc()
  go writeFunc()
  go writeFunc()

  wg.Wait()
}
