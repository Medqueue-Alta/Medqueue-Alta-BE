package middlewares

import (
	"Medqueue-Alta-BE/config"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT digunakan untuk membuat token JWT dengan ID.
func GenerateJWT(id uint) (string, error) {
    var data = jwt.MapClaims{}
    data["id"] = id
    data["iat"] = time.Now().Unix()
    data["exp"] = time.Now().Add(time.Hour * 3).Unix()

    var processToken = jwt.NewWithClaims(jwt.SigningMethodHS256, data)

    result, err := processToken.SignedString([]byte(config.JWTSECRET))

    if err != nil {
        defer func() {
            if err := recover(); err != nil {
                log.Println("error jwt creation:", err)
            }
        }()
        return "", errors.New("terjadi masalah pembuatan")
    }

    return result, nil
}


func DecodeToken(token *jwt.Token) uint {
    var result uint
    var claim = token.Claims.(jwt.MapClaims)

    if val, found := claim["id"]; found {
        result = uint(val.(float64)) 
    }

    return result
}

func DecodeTokenWithClaims(token *jwt.Token) (uint, error) {
    var result uint
    var claim = token.Claims.(jwt.MapClaims)

    if val, found := claim["id"]; found {
        result = uint(val.(float64))
    } else {
        return 0, errors.New("ID pengguna tidak ditemukan dalam token")
    }

    return result, nil
}


