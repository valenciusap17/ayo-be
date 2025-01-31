package hash

import "golang.org/x/crypto/bcrypt"

func GenerateHash(data string) (string, error) {
	bytes := []byte(data)
	hashed, err := bcrypt.GenerateFromPassword(bytes, 10)
    if err != nil {
        return data, err
    }

    return string(hashed), err
}

func ValidateHash(hashed string, data string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(data))
}