package controller

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"okr/controller/request"
	"okr/controller/response"
	"okr/domain"
	"okr/usecase"
)

type userController struct {
	uu domain.UserUsecase
}

func NewUserHandler(ur domain.UserRepository, rr domain.RedisRepository) *userController {
	return &userController{
		uu: usecase.NewUserUsecase(ur, rr),
	}
}

func (uc *userController) SignUp(c echo.Context) error {
	req := request.SignUp{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.FormBindFailed())
	}
	if err := validator.New().Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ValidateFailed())
	}

	if err := uc.uu.SignUp(req.Name, req.LoginId, req.Password, req.MailAddress); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, response.Success{Message: "登録成功"})
}

func (uc userController) SignIn(c echo.Context) error {
	req := request.SignIn{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.FormBindFailed())
	}
	if err := validator.New().Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ValidateFailed())
	}

	userToken, err := uc.uu.SignIn(req.LoginId, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, response.SignIn{Token: *userToken})
}
