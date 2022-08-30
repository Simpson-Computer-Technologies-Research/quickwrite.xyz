package main

import "github.com/realTristan/QuickWrite/server/api"

// Main Function
func main() {
	api.ListenAndServe(":8000")

	/*
		teststring := "my name is Tristan and I am pretty Cool"
		splitString := strings.Split(teststring, " ")

		st := time.Now()
		for i := 0; i < len(splitString); i++ {
			if IsAllLower(splitString[i]) && len(splitString[i]) >= 4 && len(splitString[i]) <= 15 {
				fmt.Println(splitString[i])
			}
		}
		fmt.Println(time.Since(st))
	*/
}

/* Function to check whether all the characters
// in the provided string are lowercase
func IsAllLower(w string) bool {
	for n := 0; n < len(w); n++ {
		if !IsLowerLetter(w[n]) {
			return false
		}
	}
	return true
}

// Function to check whether the provided
// byte character is a lowercase letter
func IsLowerLetter(b byte) bool {
	return b >= 97 && b <= 122
}
*/
