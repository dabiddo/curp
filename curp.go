package curp

import (
	"strconv"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

func Generate(firstName string, middleName string, paternalName string, maternalName string, dob string, gender string, myState string)string{

	//Step 1: Get First Letter and First Vowel from the Patermal Name
	//Step 2: Second letter is the first vowel from the paternal name
	//Step 3: Third letter is the first letter from maternal name
	//Step 4: Fourth letter is the first letter from the first name
	firstLetter := string(paternalName[0])
	secondLetter := firstVowel(paternalName)
	thirdLetter := string(maternalName[0])
	fourthLetter := string(firstName[0])
	alphabeticCode := firstLetter + secondLetter + thirdLetter + fourthLetter
	//fmt.Println(rune('H'))
	//For the Digit Codes, we take the las 2 digits of the Year of Birth
	// The month of birth in digit code
	//the day of birth in digit code
	dateString,_ := time.Parse(layoutISO, dob)

	year:=  strconv.Itoa(dateString.Year())

	intMonth,month := int(dateString.Month()),"0"
	// leading 0
	if intMonth < 10 {
		month += strconv.Itoa(intMonth)
	}else {
		month = strconv.Itoa(intMonth)
	}
	//leading 0
	intDay, day := dateString.Day(), "0"

	if(intDay < 10 ){
		day += strconv.Itoa(intDay)
	}else {
		day = strconv.Itoa(intDay)
	}

	birthCode := year[2:4] + month + day

	//We add the Gender H/M
	// we look for the born state
	stateKey := bornState(myState)


	generatedCurp := alphabeticCode+birthCode+gender+stateKey
	return generatedCurp
}

/**
Returns the first Vowel from the name
 */
func firstVowel(name string)string  {

	var vowelList = []rune{'A','E','I','O','U'}
	var vowel = ""
	for i,_ := range name{

		for z,_ := range vowelList{
			if string(name[i]) == string(vowelList[z]) {
				return string(name[i])
			}
		}
	}
	return vowel
}

/**
From the given state, look for the State Code string to return it.
 */
func bornState(myState string)string  {
	var states  = map[string]string{
		"BC":"BAJA CALIFORNIA",
		"MC": "MEXICO",
		"SN" : "SINALOA",
	}

	for k,v := range states  {
		if v == myState {
			return k
		}
	}
	//If its not any known state, we have to assume this person was born in another country
	return "NE"
}
