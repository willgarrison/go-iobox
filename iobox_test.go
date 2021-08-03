package iobox

import (
	"testing"
)

func TestNew(t *testing.T) {
	result := New()
	if !isIOBox(result) {
		t.Error("Expected IOBox pointer but got something else.")
	}
}

func TestConnectOutputTo(t *testing.T) {

	a := New()
	b := New()

	a.ConnectOutputTo(b)

	aOutputs := len(a.Outputs)
	bInputs := len(b.Inputs)

	if aOutputs != 1 {
		t.Error("Expected a to have 1 output, but got ", aOutputs, " outputs.")
	}
	if bInputs != 1 {
		t.Error("Expected b to have 1 input, but got ", bInputs, " inputs.")
	}
}

func TestSend(t *testing.T) {

}

func TestReceive(t *testing.T) {

}

func isIOBox(t interface{}) bool {
	switch t.(type) {
	case *IOBox:
		return true
	default:
		return false
	}
}
