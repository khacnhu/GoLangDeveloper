package unittesting

import (
	"testing"
)

func TestCheckDivisibility(t *testing.T) {

	input := 5
	want := "FIVE"
	got := CheckDivision(input)
	if want != got {
		t.Error("Incorrect Response")
	}

}
