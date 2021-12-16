package controller

import (
    "errors"

    "github.com/golang-jwt/jwt/v4"
)

func signSid(sid int) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sid": sid,
    })
    
    return token.SignedString([]byte(TestSecrkey))
}

func getSid(tokenStr string, sid *uint) error {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("Signing Method not match")
        }
        return []byte(TestSecrkey), nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        *sid = uint(claims["sid"].(float64))
    } else {
        return err
    }

    return nil
}
