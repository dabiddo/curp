package main

import (
	"fmt"
	"github.com/dabiddo/curp"
)

func main() {
	fmt.Println("***MEXICAN CURP EXAMPLE ALGORITHM***")
	fmt.Println("First Name:")
	fmt.Println("Middle Name:")
	fmt.Println("Paternal Last Name:")
	fmt.Println("Maternal Last Name:")

	//once we have all parameter, generate CURP
	myCurp := curp.Generate("LUIS","RAUL","BELLO","MENA") //TODO:get inputs from console
	//curp.IsVoewel("BELLO")
	println(myCurp)
}
