package infra

import (
	"github.com/WeCanHearYou/wechy/app"
	"github.com/WeCanHearYou/wechy/app/toolbox/env"
	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = env.MustGet("JWT_SECRET")

//Encode creates new JWT tokens with given claims
func Encode(claims *app.WechyClaims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	return jwtToken.SignedString([]byte(jwtSecret))
}

//Decode extract claims from JWT tokens
func Decode(token string) (*app.WechyClaims, error) {
	claims := &app.WechyClaims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err == nil && jwtToken.Valid {
		return claims, nil
	}

	return nil, err
}