package test

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	file, err := os.Create("test")
	if err != nil {
		t.Fatal(err)
		return
	}
	defer file.Close()

}
