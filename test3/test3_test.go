package test3

import "testing"

func TestFindFirstStringInBracket_WithEmptyString(t *testing.T) {
	actual := FindFirstStringInBracket("")
	if actual != "" {
		t.Errorf("Expected %v but got %v", "empty string", actual)
	}
}

func TestFindFirstStringInBracket_WithCorrectString(t *testing.T) {
	expectation := "test3"
	actual := FindFirstStringInBracket("(test3)")
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}

func TestFindFirstStringInBracket_WithoutOpenBracket(t *testing.T) {
	actual := FindFirstStringInBracket("test3)")
	if actual != "" {
		t.Errorf("Expected %v but got %v", "empty string", actual)
	}
}

func TestFindFirstStringInBracket_WithoutCloseBracket(t *testing.T) {
	actual := FindFirstStringInBracket("(test3")
	if actual != "" {
		t.Errorf("Expected %v but got %v", "empty string", actual)
	}
}

func TestFindFirstStringInBracket_WithoutBracket(t *testing.T) {
	actual := FindFirstStringInBracket("test3")
	if actual != "" {
		t.Errorf("Expected %v but got %v", "empty string", actual)
	}
}

func TestFindFirstStringInBracket_WithCloseBracketFirst(t *testing.T) {
	actual := FindFirstStringInBracket(")(test3")
	if actual != "" {
		t.Errorf("Expected %v but got %v", "empty string", actual)
	}
}

func TestFindFirstStringInBracket_WithDoubleBracket(t *testing.T) {
	actual := FindFirstStringInBracket("(bracket1)(bracket2))")
	expectation := "bracket1"
	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
