package utils

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

func BindRawStructToRespStruct(raw interface{}, destination interface{}) error{
	decoded, _ := json.Marshal(raw)
	err :=  json.Unmarshal(decoded, &destination)
	if err != nil {
		return err
	}
	return nil
}

func GetBodyRequest(c *gin.Context) []byte{
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	//fmt.Println("bufff: ", string(buf[0:num]))
	return buf[0:num]
}

func GetTokenFromHeader(c *gin.Context) (string, error) {
	token := c.GetHeader("Authorization")
	fmt.Println(token)
	if len(token) == 0{
		return "", errors.New("Token không khả dụng")
	}
	return token[7:], nil
}

func EncriptPwd(pwd string) string{
	hash := sha256.New()
	hash.Write([]byte(pwd))
	return fmt.Sprintf("%x", hash.Sum(nil))

}

