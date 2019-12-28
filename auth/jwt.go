package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"github.com/DoHuy/parking_to_easy/model"
	"time"
)

func hashMAC(message, key []byte) string {
	secret := "mysecret"
	data := "data"
	fmt.Printf("Secret: %s Data: %s\n", secret, data)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

type Header struct{
	HashType string
}

type Body struct{
	UserId int
	Role string
	Expired time.Time
}

type Footer struct{
	SecretKey model.SecretKey
}

type JWT struct {
	Header Header `json:"header"`
	Body Body	  `json:"body"`
	Footer Footer `json:"footer"`
}

func encrypt(hashType string, userId int, role string, expired time.Time) (encmess string, err error) {
	secretSeq := model.GetSecretKey()
	tokenComponent := JWT{
		Header: Header{
			HashType: hashType,
		},
		Body: Body{
			UserId: userId,
			Role: role,
			Expired: expired,
		},
		Footer: Footer{
			SecretKey: secretSeq,
		},
	}
	header, _ := json.Marshal(&tokenComponent.Header)
	body, 	_ := json.Marshal(&tokenComponent.Body)
	sha := hashMAC([]byte(fmt.Sprintf("%s.%s", string(header), string(body))), []byte(tokenComponent.Footer.SecretKey))
	token := fmt.Sprintf("%s.%s.%s", string(header), string(body), sha)

	// encript token
	plainText := []byte(token)
	secretKey  := []byte(tokenComponent.Footer.SecretKey)
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//returns to base64 encoded string
	encmess = base64.URLEncoding.EncodeToString(cipherText)
	return
}

func decrypt(securemess string) (decodedmess string, err error) {
	cipherText, err := base64.URLEncoding.DecodeString(securemess)
	if err != nil {
		return
	}

	secretSeq := model.GetSecretKey()
	block, err := aes.NewCipher([]byte(secretSeq))
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	decodedmess = string(cipherText)
	return
}

