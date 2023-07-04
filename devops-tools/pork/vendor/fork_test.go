package pork

import "testing"

func TestForkRepository(t *testing.T) {
	if err := ForkRepository("testing"); err != nil {
		t.Fail()
	}
}
