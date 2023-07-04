package pork

import "testing"

func TestCloneRepository(t *testing.T) {
	if err := CloneRepository("testing", "", false); err != nil {
		t.Fail()
	}
}
