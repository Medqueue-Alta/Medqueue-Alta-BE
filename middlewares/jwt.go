package middlewares

import (
	"Medqueue-Alta-BE/config"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT digunakan untuk membuat token JWT dengan ID.
func GenerateJWT(id uint, role string, nama string) (string, error) {
    var data = jwt.MapClaims{}
    data["id"] = id
    data["role"] = role
    data["nama"] =  nama
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


func DecodeToken(token *jwt.Token) (uint, string, string) {
    var userID uint
    var userRole string
    var userNama string
    var claim = token.Claims.(jwt.MapClaims)

    if val, found := claim["id"]; found {
        userID = uint(val.(float64)) 
    }

    if val, found := claim["role"];found {
        userRole = val.(string)
    }

    if val, found := claim["nama"];found {
        userNama = val.(string)
    }

    return userID, userRole, userNama
}


