package usecase

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"okr/controller/response"
	"okr/domain"
)

type userUsecase struct {
	ur domain.UserRepository
	rr domain.RedisRepository
}

func NewUserUsecase(ur domain.UserRepository, rr domain.RedisRepository) domain.UserUsecase {
	return &userUsecase{
		ur: ur,
		rr: rr,
	}
}

func (uu userUsecase) SignUp(name, loginId, password, mailAddress string) *response.ErrorResponse {

	if _, err := uu.ur.AuthByLoginId(loginId); err == nil {
		return response.AlreadyExistsData()
	}
	if _, err := uu.ur.AuthByMailAddress(mailAddress); err == nil  {
		return response.AlreadyExistsData()
	}

	user := domain.User{
		Name: name,
	}
	userId := uu.ur.InsertUser(user)
	safetyPass := createSafetyPass(password)
	auth := domain.Auth{
		UserId: userId,
		LoginId: loginId,
		Password: safetyPass,
		MailAddress: mailAddress,
	}
	uu.ur.InsertAuth(auth)

	return nil
}

func (uu userUsecase) SignIn(loginId, password string) (*string, *response.ErrorResponse) {
	auth, err := uu.ur.AuthByLoginId(loginId)
	if err != nil {
		return nil, response.NotFoundData()
	}
	if err = passwordVerify(auth.Password, password); err != nil {
		return nil, response.NotFoundData()
	}
	user, err := uu.ur.UserById(auth.UserId)
	if err != nil {
		return nil, response.NotFoundData()
	}
	userToken, _ := MakeRandomStr()
	userJson, _ := json.Marshal(user)
	if err = uu.rr.RedisSet(string(userJson), userToken); err != nil {
		return nil, response.OtherError(err)
	}

	return &userToken, nil
}

func createSafetyPass(password string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

// パスワードがハッシュにマッチするかどうかを調べる
func passwordVerify(hash, pw string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

func MakeRandomStr() (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// 乱数を生成
	b := make([]byte, 30)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error...")
	}

	// letters からランダムに取り出して文字列を生成
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}