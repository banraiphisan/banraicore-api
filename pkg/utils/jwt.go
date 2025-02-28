package utils

import (
	"strings"
	"time"

	"github.com/tubfuzzy/banraiphisan-reservation/config"
	"github.com/tubfuzzy/banraiphisan-reservation/pkg/common/exception"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

func init() {
	parser = jwt.NewParser()
}

var parser *jwt.Parser

const (
	bearerPrefix = "Bearer "
	partLength   = 2
)

func GenerateTokens(username string, roles []map[string]interface{}, config *config.Configuration) (accessToken, refreshToken string) {
	jwtSecret := config.Authorization.JWTSecret
	jwtExpired := config.Authorization.JwtExpired
	refreshTokenExpired := config.Authorization.RefreshTokenExpired

	accessClaims := jwt.MapClaims{
		"id":       uuid.New(),
		"username": username,
		"roles":    roles,
		"exp":      time.Now().Add(time.Minute * time.Duration(jwtExpired)).Unix(),
	}
	accessToken = generateTokenWithClaims(accessClaims, jwtSecret)

	refreshClaims := jwt.MapClaims{
		"id":       uuid.New(),
		"username": username,
		"exp":      time.Now().Add(time.Hour * time.Duration(refreshTokenExpired)).Unix(),
	}
	refreshToken = generateTokenWithClaims(refreshClaims, jwtSecret)

	return accessToken, refreshToken
}

func DecodeJWTFromBearerToken(bearerToken string, secret string, object interface{}) error {
	jwtToken := ExtractJWTToken(bearerToken)
	claim := jwt.MapClaims{}
	if secret != "" {
		_, err := jwt.ParseWithClaims(jwtToken, claim, func(_ *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil {
			return err
		}
	} else {
		jwtToken, _, err := parser.ParseUnverified(jwtToken, claim)
		if err != nil {
			return err
		}
		claim = jwtToken.Claims.(jwt.MapClaims)
	}

	expirationTime := time.Unix(int64(claim["exp"].(float64)), 0)
	claim["expirationTime"] = expirationTime

	return mapstructure.Decode(claim, object)
}

func generateTokenWithClaims(claims jwt.MapClaims, jwtSecret string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtSecret))
	exception.PanicLogging(err)

	return tokenSigned
}

func ExtractJWTToken(tokenString string) string {
	part := strings.Split(tokenString, bearerPrefix)
	if len(part) == partLength {
		return part[1]
	}

	return tokenString
}
