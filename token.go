package main

import (
	"fmt"
	"net/http"
	"strings"

	// Config
	"github.com/spf13/viper"

	// JWT Tokens
	"github.com/golang-jwt/jwt/v5"
)

func parseToken(tokenString string) (string, error) {
	jwtSecret := []byte(viper.GetString("JwtSecret"))
	jwtAudience := viper.GetString("JwtAudience")
	jwtRoleClaimKey := viper.GetString("JwtRoleClaimKey")

	// Parse the token from the request

	fmt.Println("jwtSecret", jwtSecret)
	fmt.Println("tokenString", tokenString)

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtSecret, nil
	}

	var token *jwt.Token
	err := error(nil)
	if jwtAudience == "" {
		token, err = jwt.Parse(tokenString, keyFunc)
	} else {
		token, err = jwt.Parse(tokenString, keyFunc, jwt.WithAudience(jwtAudience))
	}

	if err != nil {
		return "", tileAppError{
			SrcErr:   err,
			Message:  fmt.Sprintf("Failed to parse JWT token"),
			HTTPCode: http.StatusBadRequest,
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", tileAppError{
			SrcErr:   err,
			Message:  fmt.Sprintf("Failed to extract claims from JWT token"),
			HTTPCode: http.StatusBadRequest,
		}
	}

	return fmt.Sprintf("%s", claims[jwtRoleClaimKey]), nil
}

func getDatabaseRole(authHeader string) (string, error) {
	jwtSecret := viper.GetString("JwtSecret")
	anonRole := viper.GetString("AnonRole")

	// if JWT auth not configured, return empty string - queries should be run as the configured user
	if jwtSecret == "" {
		return "", nil
	}

	// if no auth header, queries should run as anon user
	if authHeader == "" {
		return anonRole, nil
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	if tokenString == "" {
		return "", tileAppError{
			//SrcErr:  err,
			Message:  fmt.Sprintf("Request included Authorization header, but this did not include a token"),
			HTTPCode: http.StatusBadRequest,
		}
	}

	// otherwise, determine which role was specified in the token
	newRole, err := parseToken(tokenString)
	return newRole, err
}

func getAuthHeader(r *http.Request) string {
	authHeader := ""
	if authorization, ok := r.Header["Authorization"]; ok {
		authHeader = authorization[0]
	}
	return authHeader
}

func getDatabaseRoleFromRequest(r *http.Request) (string, error) {
	authHeader := getAuthHeader(r)
	return getDatabaseRole(authHeader)
}
