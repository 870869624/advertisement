package user

import (
	"time"

	"github.com/jinghaijun.com/advertisement-management/db"
	"github.com/jinghaijun.com/advertisement-management/utils"
	"github.com/jinghaijun.com/advertisement-management/utils/errors"
)

type UserProlfile struct {
	Name         string `json:"nickname" gorm:"comment:用户姓名"`
	Telephone    string `json:"telephone" gorm:"index;comment:用户电话"`
	Jurisdiction string `json:"jurisdiction" gorm:"index;comment:管理机构"`
}

type UserKind int

const (
	UserKindNormal  UserKind = 1                   //普通用户
	UserKindManager UserKind = UserKindNormal << 1 //管理员
)

func (UserProlfile) TableName() string {
	return "users"
}

type Users struct {
	UserProlfile
	ID        int       `json:"ID" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"index;comment:用户账户名"`
	Password  string    `json:"password" gorm:"index;comment:用户密码"`
	CreatedAt time.Time `json:"created_at"`
}

//加密密码
func (u *Users) Encrypt_Password() {
	u.Password = utils.Encrypt(u.Password)
}

//检验数据是否正确
func (u *Users) Validate() *errors.Error {
	if u.Username == "" || u.Password == "" {
		e := errors.New("参数错误")
		return e
	}
	return nil
}

//检验用户是否存在
func (u *Users) DoesSameExists() bool {
	var count int64
	connection := db.Get_DB()
	connection.Table("users").Where("username = ?", u.Username).Count(&count)
	return count > 0
}
func (users *Users) Create() error {
	connection := db.Get_DB()
	if err := users.Validate(); err != nil {
		return err
	}
	if users.DoesSameExists() {
		return errors.New("数据已经存在")
	}
	users.Encrypt_Password()
	err := connection.Create(users)
	if err.Error != nil {
		e := errors.New("参数有误")
		return e
	}
	return nil
}

func (users *Users) Update() error {
	var user Users
	connection := db.Get_DB()
	find := connection.Find(&user)
	if find.Error != nil {
		return find.Error
	}
	result := find.Updates(users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (user *Users) Delete() {

}
