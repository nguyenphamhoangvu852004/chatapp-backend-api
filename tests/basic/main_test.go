package basic

import "testing"

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	acctual := AddOne(1)
	if acctual != output {
		t.Errorf("AddOne(%d) = %d, want %d", input, acctual, output)
	}
}
