package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func parseToken(token *jwt.Token) (interface{}, error) {
	secret := []byte("")
	return secret, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s Your.JWT.Here\n", os.Args[0])
		os.Exit(1)
	}

	tokenStr := os.Args[1]
	if token, err := jwt.Parse(tokenStr, parseToken); err != nil && err.Error() != "signature is invalid" {
		fmt.Println(err)
		os.Exit(1)
	} else {
		b, _ := json.MarshalIndent(token.Claims, "", " ")
		fmt.Println(string(b))
	}
}
