package indexer

import (
	"testing"
)

func TestRead(t *testing.T) {

	sd, err := Read("../test/input")
	if err != nil {
		t.Log("Read error")
		t.Fail()
	}

	if len(sd.Slides) < 1 {
		t.Log("Empty slides error")
		t.Fail()
	}

}
