package shipment

import "testing"

func TestExtractOrderIDFromMsg(t *testing.T) {
	msg := "1234 paid"
	actual, err := ExtractOrderIDFromMsg(msg)

	expected := uint32(1234)

	if err != nil {
		t.Errorf("error")
	}

	if expected != actual {
		t.Errorf("Expected %d, but was %d", expected, actual)
	}
}
