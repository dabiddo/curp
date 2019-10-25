package curp

import (
	"testing"
)

func TestCurp(t *testing.T) {

	myCurp := Generate("LUIS", "RAUL", "BELLO", "MENA")

	if  myCurp != "BEML"{
		t.Error("CURP is Incorrect",myCurp)
	}
}
