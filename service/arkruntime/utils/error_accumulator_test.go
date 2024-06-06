package utils

import (
	"bytes"
	"testing"
)

func TestErrorAccumulatorBytes(t *testing.T) {
	accumulator := &DefaultErrorAccumulator{
		Buffer: &bytes.Buffer{},
	}

	errBytes := accumulator.Bytes()
	if len(errBytes) != 0 {
		t.Fatalf("Did not return nil with empty bytes: %s", string(errBytes))
	}

	err := accumulator.Write([]byte("{}"))
	if err != nil {
		t.Fatalf("%+v", err)
	}

	errBytes = accumulator.Bytes()
	if len(errBytes) == 0 {
		t.Fatalf("Did not return error bytes when has error: %s", string(errBytes))
	}
}
