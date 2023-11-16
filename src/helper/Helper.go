package helper

import (
	"crypto/md5"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	User_id  uint   `json:"identity"`
	Username string `json:"name"`
	Userrole string `json:"is_admin"`
	jwt.StandardClaims
}

var myKey = []byte("EAMS")

func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
func GenerateToken(user_id uint, username, userrole string) (string, error) {
	UserClaim := &UserClaims{
		User_id:        user_id,
		Username:       username,
		Userrole:       userrole,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}
