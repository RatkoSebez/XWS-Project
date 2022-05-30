package utilities

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

const Expiration = 86400

var TokenSecretKey = ""

type TokenInnerData struct {
	AgentName string `json:"agentName"`
	jwt.StandardClaims
}

func initAgentToken() {
	env := os.Getenv("PUBLIC_JWT_TOKEN_SECRET")
	if env == "" {
		TokenSecretKey = "agent_secret"
	} else {
		TokenSecretKey = env
	}
}

func CreateAgentToken(agentName string, issuer string) (string, error) {
	if TokenSecret == "" {
		initAgentToken()
	}
	claims := TokenInnerData{AgentName: agentName, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + ExpiresIn,
		IssuedAt:  time.Now().Unix(),
		Issuer:    issuer,
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(TokenSecretKey))
}

func GetAgentNameFromPureToken(tok string) string {
	if TokenSecretKey == "" {
		initPublicToken()
	}
	token, err := jwt.ParseWithClaims(tok, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(TokenSecretKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if claims, ok := token.Claims.(*TokenInnerData); ok && token.Valid {
		return claims.AgentName
	} else {
		fmt.Println(err)
		return ""
	}
}
