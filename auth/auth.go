package auth

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/config"
	"github.com/DoHuy/parking_to_easy/model"
	"time"
)

func CheckTokenIsTrue(authHeader string) (bool, error) {
	secretKey  := string(config.GetSecretKey())
	_, err := Decode(authHeader[7:], secretKey)
	if err != nil {
		return false, fmt.Errorf("%s", err.Error())
	}
	return true, nil
}

func CheckExpiredToken(token string) (bool, error) {
	secretKey    := string(config.GetSecretKey())
	var payload model.Payload
	decoded, err := Decode(token, secretKey)
	json.Unmarshal(decoded, &payload)
	expired, err := time.Parse(time.RFC3339, payload.Expired)
	if err != nil {
		return false, fmt.Errorf("Lỗi trên hệ thống: %s", err.Error())
	}

	if time.Now().Unix() - expired.Unix() > config.GetTokenExpired() {
		return false, nil
	}

	return true, nil

}
