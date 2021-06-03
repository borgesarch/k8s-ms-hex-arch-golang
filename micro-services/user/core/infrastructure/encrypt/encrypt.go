package encrypt

import (
	"encoding/base64"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type IEncrityHash interface {
	HashPassword(password string) string
	DoPasswordsMatch(hashedPassword, currPassword string) bool
}

type EncrityHash int

func (c EncrityHash) HashPassword(password string) string {

	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedPasswordBytes)
	return base64EncodedPasswordHash
}

func (c EncrityHash) DoPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currPassword))
	return err != nil
}
