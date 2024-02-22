package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jinghaijun.com/advertisement-management/models/user"
)

const (
	VERIFY_KEY = "advertisement"
)

type UserAuthorazation struct {
	username string
	Kind     user.UserKind
}
type UserClaims struct {
	*jwt.RegisteredClaims
	*UserAuthorazation
}

func (c *UserAuthorazation) Create_JWt() (string, error) {
	claim := &UserClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		c,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return t.SignedString([]byte(VERIFY_KEY))
}

func New(username string, kind user.UserKind) (string, error) {
	user := UserAuthorazation{
		username: username,
		Kind:     kind,
	}
	return user.Create_JWt()
}
func Parse(tokenString string) (*UserAuthorazation, error) {
	token, e := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(VERIFY_KEY), nil
	})
	if !token.Valid {
		return nil, e
	}
	claim := token.Claims.(*UserClaims)
	return claim.UserAuthorazation, nil
}
