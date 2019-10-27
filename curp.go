package curp

func Generate(firstName string, middleName string, paternalName string, maternalName string, yearOfBirth string, monthOfBirth string, dayOfBirth string)string{

	//Step 1: Get First Letter and First Vowel from the Patermal Name
	//Step 2: Second letter is the first vowel from the paternal name
	//Step 3: Third letter is the first letter from maternal name
	//Step 4: Fourth letter is the first letter from the first name
	firstLetter := string(paternalName[0])
	secondLetter := IsVoewel(paternalName)
	thirdLetter := string(maternalName[0])
	fourthLetter := string(firstName[0])
	alphabeticCode := firstLetter + secondLetter + thirdLetter + fourthLetter

	//For the Digit Codes, we take the las 2 digits of the Year of Birth
	// The month of birth in digit code
	//the day of birth in digit code
	year := string(yearOfBirth[2])+string(yearOfBirth[3])
	birthCode := year + monthOfBirth + dayOfBirth

	generatedCurp := alphabeticCode+birthCode
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
