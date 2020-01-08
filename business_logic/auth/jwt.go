package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DoHuy/parking_to_easy/model"
	"strings"
)


func CompareHmac(signatureValue, token, secret string) bool {
	return isValidHash(signatureValue, token, secret)
}

// Base64Encode takes in a string and returns a base 64 encoded string
func Base64Encode(src string) string {
	return strings.
		TrimRight(base64.URLEncoding.
			EncodeToString([]byte(src)), "=")
}

// Base64Encode takes in a base 64 encoded string and returns the //actual string or an error of it fails to decode the string
func Base64Decode(src string) (string, error) {
	if l := len(src) % 4; l > 0 {
		src += strings.Repeat("=", 4-l)
	}
	decoded, err := base64.URLEncoding.DecodeString(src)
	fmt.Sprintln("base64 decode")
	if err != nil {
		errMsg := fmt.Errorf("Decoding Error %s", err)
		return "", errMsg
	}
	return string(decoded), nil
}

// Hash generates a Hmac256 hash of a string using a secret
func Hash(src string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(src))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// isValidHash validates a hash againt a value
func isValidHash(value string, hash string, secret string) bool {
	return hash == Hash(value, secret)
}

func Encode(payload model.Payload, secret string) (string, error) {
	header := model.Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	//var str []byte
	str, _ := json.Marshal(header)
	headerToBase64 := Base64Encode(string(str))
	encodedPayload, _ := json.Marshal(payload)
	signatureValue := fmt.Sprintf("%s.%s", headerToBase64, Base64Encode(string(encodedPayload)))
	return fmt.Sprintf("%s.%s", signatureValue, Hash(signatureValue, secret)), nil
}

func Decode(jwt string, secret string) ([]byte, error) {
	token := strings.Split(jwt, ".")
	// check if the jwt token contains
	// header, payload and token
	if len(token) != 3 {
		splitErr := errors.New("Invalid token: token should contain header, payload and secret")
		return nil, splitErr
	}
	// decode payload
	decodedPayload, PayloadErr := Base64Decode(token[1])
	if PayloadErr != nil {
		return nil, fmt.Errorf("Invalid payload: %s", PayloadErr.Error())
	}
	payload := model.Payload{}
	// parses payload from string to a struct
	ParseErr := json.Unmarshal([]byte(decodedPayload), &payload)
	if ParseErr != nil {
		return nil, fmt.Errorf("Invalid payload: %s", ParseErr.Error())
	}
	signatureValue := token[0] + "." + token[1]
	// verifies if the header and signature is exactly whats in
	// the signature
	if CompareHmac(signatureValue, token[2], secret) == false {
		return nil, errors.New("Invalid token")
	}
	result, _ := json.Marshal(payload)
	return result, nil
}
