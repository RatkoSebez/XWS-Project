package utilities

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
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
		TokenSecretKey = "token_secret"
	} else {
		TokenSecretKey = env
	}
}

func CreateAgentToken(agentName string, issuer string) (string, error) {
	if TokenSecretKey == "" {
		initAgentToken()
	}
	claims := TokenInnerData{AgentName: agentName, StandardClaims: jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + Expiration,
		IssuedAt:  time.Now().Unix(),
		Issuer:    issuer,
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(TokenSecretKey))
}

func GetAgentNameFromPureToken(tok string) string {
	if TokenSecretKey == "" {
		initAgentToken()
	}
	token, err := jwt.ParseWithClaims(tok, &TokenInnerData{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(TokenSecretKey), nil
	})
	if err != nil {
		fmt.Println(err)
	}
	if claims, ok := token.Claims.(*TokenInnerData); ok && token.Valid {
		return claims.AgentName
	} else {
		fmt.Println(err)
		return ""
	}
}

func getAgentToken(header http.Header) (string, error) {
	reqToken := header.Get("AgentToken")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return "", fmt.Errorf("NO_TOKEN")
	}
	return strings.TrimSpace(splitToken[1]), nil
}

func GetAgentNameFromToken(r *http.Request) string {
	if TokenSecretKey == "" {
		initAgentToken()
	}
	tokenString, err := getAgentToken(r.Header)
	if err != nil {
		panic(err)
	}
	return GetAgentNameFromPureToken(tokenString)
}
