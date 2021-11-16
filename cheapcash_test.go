package cheapcash_test

import (
	"log"
	"os"
	"testing"

	"github.com/aldy505/cheapcash"
)

func TestMain(m *testing.M) {
	removeDirIfExists("/tmp/cheapcash")
	defer removeDirIfExists("/tmp/cheapcash")

	os.Exit(m.Run())
}

func TestDefault(t *testing.T) {
	c := cheapcash.Default()
	if c.Path != "/tmp/cheapcash/" {
		t.Error("expected path to return /tmp/cheapcash/, got:", c.Path)
	}
}

func TestNew(t *testing.T) {
	c := cheapcash.New("/somewhere")
	if c.Path != "/somewhere" {
		t.Error("expected path to return /somewhere, got:", c.Path)
	}
}

func removeDirIfExists(path string) {
	dir, err := os.Stat(path)
	if err == nil {
		if dir.IsDir() {
			err = os.RemoveAll(path)
			if err != nil {
				log.Fatal("unable to remove temp directory:", err)
			}
		}
	}
}
