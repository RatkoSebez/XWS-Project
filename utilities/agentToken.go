package utilities

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

const ExpiresIn = 86400

var TokenSecret = ""

type TokenClaims struct {
	AgentName string `json:"agentName"`
	jwt.StandardClaims
}

func initAgentToken() {
	env := os.Getenv("PUBLIC_JWT_TOKEN_SECRET")
	if env == "" {
		TokenSecret = "agent_secret"
	} else {
		TokenSecret = env
	}
}

func CreateToken(agentName string, issuer string) (string, error) {
	if TokenSecret == "" {
		initAgentToken()
	}
	claims := TokenClaims{AgentName: agentName, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + ExpiresIn,
		IssuedAt:  time.Now().Unix(),
		Issuer:    issuer,
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(TokenSecret))
}

func GetAgentNameFromPureToken(tok string) string {
	if TokenSecret == "" {
		initPublicToken()
	}
	token, err := jwt.ParseWithClaims(tok, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(TokenSecret), nil
	})
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims.AgentName
	} else {
		fmt.Println(err)
		return ""
	}
}
