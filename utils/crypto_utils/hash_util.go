package crypto_utils

import "golang.org/x/crypto/bcrypt"

// HashUtils holds the methods of hashed utils
type HashUtils interface {
	GenerateBcrtptHash(text string) (string, error)
	MatchBcryptedHash(hashedText, plainText string) error
}

type hashUtils struct {
}

//NewHashUtils implements HashUtils
func NewHashUtils() HashUtils {
	return hashUtils{}
}

//GenerateBcrtptHash is the function that generates hash og given text
func (hashUtils hashUtils) GenerateBcrtptHash(text string) (string, error) {
	//create password hash from original password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(passwordHash), nil
}

// MatchBcryptedHash compair hased text and paintext text
func (hashUtils hashUtils) MatchBcryptedHash(hashedText, plainText string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))
}
