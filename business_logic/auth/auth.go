package auth

import (
	"encoding/json"
	"fmt"
	"github.com/DoHuy/parking_to_easy/config"
	"github.com/DoHuy/parking_to_easy/model"
	"github.com/DoHuy/parking_to_easy/mysql"
	redis2 "github.com/DoHuy/parking_to_easy/redis"
	"time"
)

type Auth struct {
	Redis *redis2.Redis
}

// singleton init a instance redis
func NewAuth() *Auth {
	redis := redis2.NewRedis()
	return &Auth{Redis: redis,}
}

func (this *Auth) CheckTokenIsTrue(token string) (bool, error) {
	secretKey := string(config.GetSecretKey())
	_, err := Decode(token, secretKey)
	if err != nil {
		return false, fmt.Errorf("%s", err.Error())
	}
	return true, nil
}

func (this *Auth) CheckExpiredToken(token string) (bool, error, error) { //done neu expired thi return true
	secretKey := string(config.GetSecretKey())
	var payload model.Payload
	decoded, err := Decode(token, secretKey)
	json.Unmarshal(decoded, &payload)
	expired, err := time.Parse(time.RFC3339, payload.Expired)
	if err != nil {
		return true, nil, fmt.Errorf("Lỗi trên hệ thống: %s", err.Error())
	}
	if expired.UnixNano()/1000000-time.Now().UnixNano()/1000000 >= 0 {
		return false, nil, nil
	}

	return true, nil, nil

}

// authenticate user login return token and error
func (this *Auth) Authenticate(credential model.Credential, credIFace mysql.CredentialDAO) (string, error, error) {
	var err error
	credential, err = credIFace.FindCredentialByNameAndPassword(credential.Username, credential.Password)
	if err != nil {
		return "", fmt.Errorf("username hoặc password không đúng"), nil
	}
	// lay secret key cho viec giai ma token
	secretKey := string(config.GetSecretKey())
	jwtModel, err := this.Redis.GetJWTTokenFromRedis(credential.Username)
	fmt.Println("asdadasd roleeeee", credential.Role)
	// Neu user chua co token, Tạo mới token lưu vào redis và mysql
	expiredDuration := time.Minute * time.Duration(config.GetTokenExpired())
	if len(jwtModel.Jwt) <= 0 {
		// hard code
		//fmt.Println("expiredDuration:   ", expiredDuration.Milliseconds())
		jwt, err := Encode(model.Payload{UserId: credential.ID, Role: credential.Role, Expired: time.Now().Add(expiredDuration).Format(time.RFC3339),}, secretKey)
		err = this.Redis.SetJWTTokenToRedis(credential.Username, jwt)
		if err != nil {
			return "", nil, fmt.Errorf("Error with Redis %s", err.Error())
		}
		return jwt, nil, nil
	} else { // neu user daco token thi kiem tra han cua token
		decodedData, err := Decode(jwtModel.Jwt, secretKey)
		fmt.Println("deaddededed", string(decodedData))
		if err != nil {
			return "", nil, fmt.Errorf("Loi nay xay ra khi giai ma token: %s", err.Error())
		}
		var payload model.Payload
		//fmt.Println("Co caci gi do dasdadadsd")
		err = json.Unmarshal(decodedData, &payload)
		if err != nil {
			return "", nil, fmt.Errorf("Loi xay ra khi parser payload: %s", err.Error())
		}

		fmt.Println("PAYLOAD: ", payload)
		//fmt.Println("asdadasd", jwtModel)
		checked, err, _ := this.CheckExpiredToken(jwtModel.Jwt)
		fmt.Println("checked han:::::", checked)
		if checked {
			fmt.Println("checked han::::: co tao moi", checked)
			// tao moi token len redis va tra e token moi
			newJwt, err := Encode(model.Payload{UserId: credential.ID, Role: credential.Role, Expired: time.Now().Add(expiredDuration).Format(time.RFC3339),}, secretKey)
			err = this.Redis.DelJWTToken(credential.Username)
			err = this.Redis.SetJWTTokenToRedis(credential.Username, newJwt)
			if err != nil {
				return "", nil, fmt.Errorf("Loi xay ra khi lam viec voi Redis: %s", err.Error())
			}
			return newJwt, nil, nil
		}
		//
		//c.JSON(http.StatusOK, map[string]string{"token": jwt.Jwt})
		return jwtModel.Jwt, nil, nil
	}
}

// accept or reject api calling
func (this *Auth) Authorize(params ...string) (string, error, error) {
	secret := string(config.GetSecretKey())
	decoded, err := Decode(params[0], secret)
	if err != nil {
		return "", fmt.Errorf("Hệ thống có sự cố"), err
	}
	var payload model.Payload
	err = json.Unmarshal(decoded, &payload)

	switch len(params) {

	}
	if err != nil {
		return "", fmt.Errorf("Hệ thống có sự cố"), err
	}
	return payload.Role, nil, nil
}
