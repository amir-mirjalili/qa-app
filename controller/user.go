package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qa-app/pkg/httpmsg"
	"qa-app/request/user"
	"qa-app/service"
	userValidator "qa-app/validator/user-validator"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (controller *UserController) Create(ctx *gin.Context) {
	createUserRequest := user.CreateUserRequest{}
	err := ctx.ShouldBind(&createUserRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if fieldErrors, err := userValidator.UserValidateRegisterRequest(createUserRequest); err != nil {
		msg, code := httpmsg.Error(err)
		ctx.JSON(code, gin.H{"message": msg, "errors": fieldErrors})
		return
	}

	controller.userService.Create(createUserRequest)
	ctx.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	userResponse := controller.userService.FindAll()
	ctx.JSON(http.StatusOK, gin.H{"data": userResponse})
}
