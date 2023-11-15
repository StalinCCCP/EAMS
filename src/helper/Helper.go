package helper

import (
	"crypto/md5"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	user_id  uint   `json:"identity"`
	username string `json:"name"`
	userrole string `json:"is_admin"`
	jwt.StandardClaims
}

var myKey = []byte("EAMS")

func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
func GenerateToken(user_id uint, username, userrole string) (string, error) {
	UserClaim := &UserClaims{
		user_id:        user_id,
		username:       username,
		userrole:       userrole,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
