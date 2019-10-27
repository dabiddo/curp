package curp

import (
	"testing"
)

func TestCurp(t *testing.T) {

	myCurp := Generate("LUIS", "RAUL", "BELLO", "MENA", "1992","03","13")

	if  myCurp != "BEML920313"{
		t.Error("CURP is Incorrect",myCurp)
	}
}
