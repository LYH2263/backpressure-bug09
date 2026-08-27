package backpressure_test

import (
	"testing"

	"github.com/LYH2263/go-backpressure"
)

func newEngine(t *testing.T) *backpressure.Engine {
	e, err := backpressure.New(backpressure.Options{})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = e.Close() })
	return e
}
