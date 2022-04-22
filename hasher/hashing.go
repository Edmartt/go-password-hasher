//This package provide a function to convert plain text to hash
package hasher

import (
	"log"

	"golang.org/x/crypto/bcrypt" //For generating random salted hash from text plain
)

//This function takes a string and converts it to hash with pbdkf2 and sha256 based on bcrypt package
func ConvertToHash(password string) string {
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}
