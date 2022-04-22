# Password Hasher

This project is intended to ease the use of **bcrypt** package abstracting the developer from using arrays of bytes and all necessary stuff for hashing a string and making it as simple as just to call the need function and passing a string as argument

### What you can do:
- convert a plaintext string to hash 
- checks if a plaintext string is equal to a hash


#### How to use

##### hashing strings

```
import (
	"fmt"
	"github.com/Edmartt/go-password-hasher/hasher"

	)

func main(){

	myString := "12345"
	hashedString := hasher.ConvertToHash(myString)

	fmt.Println(hashedString)

}
```

We'll see some output like this:
- **$2a$10$J65kHWsGVPLJq47D5aBJmeBpytRtXM5F6iN4ZId/Eum5IXw4cOMfi**


##### Checking if a hash equals to a string we pass as argument:

```
import (
	"fmt"
	"github.com/Edmartt/go-password-hasher/hasher"

	)

func main(){

	myString := "12345"
	hashedString := hasher.ConvertToHash(myString)

	isTheSame := hasher.CheckHash(hashedString, "12345")

	fmt.Println(isTheSame)
}

```
Our output here would be:
- **true
