package curp

func Generate(firstName string, middleName string, paternalName string, maternalName string)string{

	//Step 1: Get First Letter and First Vowel from the Patermal Name
	//Step 2: Second letter is the first vowel from the paternal name
	//Step 3: Third letter is the first letter from maternal name
	//Step 4: Fourth letter is the first letter from the first name
	firstLetter := string(paternalName[0])
	secondLetter := IsVoewel(paternalName)
	thirdLetter := string(maternalName[0])
	fourthLetter := string(firstName[0])

	generatedCurp := firstLetter+secondLetter+thirdLetter+fourthLetter

	return generatedCurp
}

func IsVoewel(name string)string  {

	var vowelList = []rune{'A','E','I','O','U'}
	var vowel = ""
	for i,_ := range name{

		for z,_ := range vowelList{
			if string(name[i]) == string(vowelList[z]) {
				//println(string(name[i]))
				return string(name[i])
			}
		}
	}
	return vowel
}
