//This package provide a function to convert plain text to hash
package hasher

import "golang.org/x/crypto/bcrypt"

//This function checks if the already hashed string equals to the new plain text string
func CheckHash(currentPasswordHash string, password string) bool{
	isTheSame := bcrypt.CompareHashAndPassword([]byte(currentPasswordHash), []byte(password))


	if isTheSame != nil{
		return false
	}
	return true
}
