package token

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

//TODO move this const to config file
const RememberTokenBytes = 64

//generateRandomString used to generate a random string

func generateRandomString(n uint) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", nil
	}
	return base64.URLEncoding.EncodeToString(b), err

}

//GenerateToken used to generate a random string  used
// as token
func GenerateToken(Nbytes uint) (string, error) {
	str, err := generateRandomString(Nbytes)
	if err != nil {
		return "", err
	}
	return str, nil
}

// HashToken is used to hash a strings with a provided key string it adds the  key to the random generated string
//and  hash it using the sha256() and then return the hash as string
func HashToken(input, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Reset()
	h.Write([]byte(input))
	b := h.Sum(nil)
	return base64.URLEncoding.EncodeToString(b)

}