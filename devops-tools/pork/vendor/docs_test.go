package pork

import "testing"

func TestGetRepositoryDocs(t *testing.T) {
	content := GetRepositoryDoc("testing")
	if content != "testing" {
		t.Fail()
	}
}
