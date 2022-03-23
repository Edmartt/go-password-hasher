package hasher

import (
	"golang.org/x/crypto/bcrypt" //For generating random salted hash from text plain
)

//Takes casted to byte string and hash it with a random salt each time
func HashMaker(password []byte) string {
	hashed_password, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed_password)
}
