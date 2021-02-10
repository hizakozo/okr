package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"okr/controller"
	"okr/repository"
)

var Echo *echo.Echo

func init()  {
	e := echo.New()
	e.Use(middleware.CORS())

	ur := repository.NewUserRepository(Db)
	rr := repository.NewRedisRepository(Redis)

	userController := controller.NewUserHandler(ur, rr)
	e.POST("/user/signup", func(c echo.Context) error { return userController.SignUp(c) })
	e.POST("/user/signin", func(c echo.Context) error { return userController.SignIn(c) })
	Echo = e
}

func Run()  {
	Echo.Logger.Fatal(Echo.Start(":1313"))
}