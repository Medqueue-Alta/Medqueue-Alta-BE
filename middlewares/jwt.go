package middlewares

import (
	"Medqueue-Alta-BE/config"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type mdJwt struct{}

type JwtInterface interface {
	GenerateJWT(id uint) (string, error)
	DecodeToken(token *jwt.Token) uint
	GetUserID(r *http.Request) (uint, error)
}

func NewMidlewareJWT() JwtInterface {
	return &mdJwt{}
}

// GenerateJWT digunakan untuk membuat token JWT dengan ID.
func (md *mdJwt) GenerateJWT(id uint) (string, error) {
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

func (md *mdJwt) DecodeToken(token *jwt.Token) uint {
	var result uint
	var claim = token.Claims.(jwt.MapClaims)

	if val, found := claim["id"]; found {
		result = uint(val.(float64))
	}

	return result
}

func (md *mdJwt) GetUserID(r *http.Request) (uint, error) {
	// Get the authorization header from the request
	tokenString := r.Header.Get("Authorization")

	// Check if the authorization header is present
	if tokenString == "" {
		return 0, errors.New("authorization header is missing")
	}

	// Remove the "Bearer " prefix from the token string
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Parse the JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// Return the key used to sign the token
		return []byte(config.JWTSECRET), nil
	})
	if err != nil {
		return 0, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract the user ID from the claims
		userID := uint(claims["id"].(float64))
		return userID, nil
	}

	return 0, errors.New("invalid token")
}
