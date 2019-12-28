package auth

func CheckTokenIsTrue(encriptMsg string) bool {
	token, _ := decrypt(encriptMsg)

	return false
}

func CheckExpiredToken(token string) bool {

}
