package authentication

import (
	"errors"

	"github.com/jinghaijun.com/advertisement-management/db"
	"github.com/jinghaijun.com/advertisement-management/models/token"
	"github.com/jinghaijun.com/advertisement-management/models/user"
)

type Authentication struct {
	Username string
	Password string `gorm:"index"`
}

func (a *Authentication) Create_JWt() (string, error) {
	return token.New(a.Username, user.UserKindNormal)
}
func (a *Authentication) Signin() (string, error) {
	u := &user.Users{
		Username: a.Username,
		Password: a.Password,
	}
	if err := u.Validate(); err != nil {
		return "", err
	}
	if !u.DoesSameExists() {
		return "", errors.New("用户不存在")
	}
	u.Encrypt_Password()
	connection := db.Get_DB()
	e := connection.Table("users").Where("username = ? and password = ?", u.Username, u.Password).First(&a)
	if e.Error != nil {
		return "", e.Error
	}
	return a.Create_JWt()
}
