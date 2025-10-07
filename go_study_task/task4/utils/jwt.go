package utils

import (
  "fmt"
  "github.com/golang-jwt/jwt"
  "log"
  "time"
)

var jwtSecret = []byte("zhangxuehai-secret")

func GenerateToken(userID uint, roles []string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserID": userID,
		"roles":  roles,
		"exp":    time.Now().Add(8 * time.Hour).Unix(),
	})
	token1, err := token.SignedString(jwtSecret)
	if err != nil {
	  log.Printf("err: %s", err.Error())
  }
  return token1, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("unexpected signing method:%v\n", token.Header["alg"])
    }
    return jwtSecret, nil
  })
  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    return claims, nil
  } else {
    return nil, err
  }
}
