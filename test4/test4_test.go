package test4

import (
	"reflect"
	"testing"
)

func TestGroupWordsByAnagram_WithEmptySlice(t *testing.T) {
	actual := GroupWordsByAnagram([]string{})
	if actual != nil {
		t.Errorf("Expected %v but got %v", nil, actual)
	}
}

func TestGroupWordsByAnagram_WithCorrectSlice(t *testing.T) {
	expectation := map[string][]string{
		"aakmn": {"makan"},
		"aik":   {"kia"},
		"aikt":  {"kita", "atik", "tika"},
		"aku":   {"aku", "kua"},
	}
	words := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	actual := GroupWordsByAnagram(words)
	if !reflect.DeepEqual(actual, expectation) {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
