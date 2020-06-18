package order

import "testing"

func TestExtractOrderIDFromMsg(t *testing.T) {
	msg := "1234 payed"
	actual := ExtractEventMsg(msg)

	expected := "payed"

	if expected != actual {
		t.Errorf("Expected %v, but was %v", expected, actual)
	}
}
