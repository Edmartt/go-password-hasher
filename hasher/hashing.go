package hasher

import (
	"golang.org/x/crypto/bcrypt" //For generating random salted hash from text plain
)

//This function takes a string and converts it to hash with pbdkf2 and sha256
func ConvertToHash(password string) string {
	
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed_password)
}
