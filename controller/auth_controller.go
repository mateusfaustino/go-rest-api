// controller/auth_controller.go
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusfaustino/go-rest-api/usecase"
)

type AuthController struct {
	AuthUseCase usecase.AuthUseCase
}

func NewAuthController(authUseCase usecase.AuthUseCase) AuthController {
	return AuthController{
		AuthUseCase: authUseCase,
	}
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := ac.AuthUseCase.Login(loginData.Username, loginData.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
