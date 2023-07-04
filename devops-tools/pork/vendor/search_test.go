package pork

import "testing"

func TestSearchKeywords(t *testing.T) {
	repositoryList := SearchByKeyword([]string{"one", "two"})
	if repositoryList[0] != "myrepository" {
		t.Fail()
	}
}
