package repository

import (
	"github.com/jinzhu/gorm"
	"okr/domain"
)

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) domain.UserRepository {
	return &userRepository{
		Db: Db,
	}
}

func (ur userRepository) InsertUser(user domain.User) int {
	ur.Db.Create(&user)
	return user.ID
}

func (ur userRepository) InsertAuth(auth domain.Auth) {
	ur.Db.Create(&auth)
}

func (ur userRepository) AuthByLoginId(loginId string) (*domain.Auth, error) {
	auth := domain.Auth{}
	err := ur.Db.Select("auth_id, user_id, login_id, password, mail_address").
		Table("auth").
		Where("login_id = ?", loginId).
		Find(&auth).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	return &auth, err
}

func (ur userRepository) UserById(userId int) (*domain.User, error) {
	user := domain.User{}
	err := ur.Db.Select("user_id, user_name").
		Table("user").
		Where("user_id  = ?", userId).
		Find(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	return &user, err
}

func (ur userRepository) AuthByMailAddress(mailAddress string) (*domain.Auth, error) {
	auth := domain.Auth{}
	err := ur.Db.Select("auth_id, user_id, login_id, password, mail_address").
		Table("auth").
		Where("mail_address = ?", mailAddress).
		Find(&auth).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, err
	}
	return &auth, err
}
