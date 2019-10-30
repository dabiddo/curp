package curp

import (
	"testing"
)

func TestCurp(t *testing.T) {

	myCurp := Generate("LUIS", "RAUL", "BELLO", "MENA", "1992-03-13","H","MEXICO")

	if  myCurp != "BEML920313HMC"{
		t.Error("CURP is Incorrect",myCurp)
	}
}
