package domain

import "okr/controller/response"

type User struct {
	ID       int    `gorm:"column:user_id;PRIMARY_KEY"`
	Name     string `gorm:"column:user_name"`
	IsDelete int `gorm:"column:is_delete;default:'galeone'"`
}

type Auth struct {
	ID          int    `gorm:"column:auth_id;PRIMARY_KEY"`
	UserId      int    `gorm:"column:user_id"`
	LoginId     string `gorm:"column:login_id"`
	Password    string `gorm:"column:password"`
	MailAddress string
}

type UserUsecase interface {
	SignUp(name, loginId, password, mailAddress string) *response.ErrorResponse
	SignIn(loginId, password string) (*string, *response.ErrorResponse)
}

type UserRepository interface {
	InsertUser(user User) int
	InsertAuth(auth Auth)
	UserById(userId int) (*User, error)
	AuthByLoginId(loginId string) (*Auth, error)
	AuthByMailAddress(mailAddress string) (*Auth,error)
}

type RedisRepository interface {
	RedisSet(json string, key string) error
	RedisGet(key string) (*User, error)
	RedisDelete(token string)
}