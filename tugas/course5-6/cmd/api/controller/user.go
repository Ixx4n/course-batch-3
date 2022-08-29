package controller

import (
	"course5-6/cmd/model"
	"course5-6/cmd/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	userSvc *service.UserSvc
}

func NewUserController(DB *gorm.DB) *UserController {
	userSvc := service.NewUserSvc(DB)
	return &UserController{userSvc}
}

func (userController UserController) Register(c *gin.Context) {
	var registerRequest model.RegisterRequest
	if err := c.ShouldBind(&registerRequest); err != nil {
		Handler(c, 400, err)
		return
	}
	code, token, err := userController.userSvc.Register(registerRequest)
	if err == nil {
		c.JSON(code, map[string]string{
			"token": token,
		})
	} else {
		Handler(c, code, err)
	}
}

func (userController UserController) Login(c *gin.Context) {
	var loginRequest model.LoginRequest
	if err := c.ShouldBind(&loginRequest); err != nil {
		Handler(c, 400, err)
		return
	}
	code, token, err := userController.userSvc.Login(loginRequest)
	if err == nil {
		c.JSON(code, map[string]string{
			"token": token,
		})
	} else {
		Handler(c, code, err)
	}
}
