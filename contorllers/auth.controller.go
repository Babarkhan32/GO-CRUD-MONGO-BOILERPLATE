package controllers

import (
	"net/http"

	"exmaple.com/job-x-apis/helpers"
	"exmaple.com/job-x-apis/models"
	"exmaple.com/job-x-apis/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuth(authService services.AuthService) AuthController {
	return AuthController{
		AuthService: authService,
	}
}

func (uc *AuthController) Login(ctx *gin.Context) {
	var credentials models.Credentials
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	user, err := uc.AuthService.Login(&credentials)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	token, err := helpers.CreateJWT(user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}

	ctx.JSON(http.StatusCreated, gin.H{"data": user, "token": token})

}

func (uc *AuthController) CheckStatus(ctx *gin.Context) {
	message := uc.AuthService.CheckStatus()
	ctx.JSON(http.StatusOK, gin.H{"message": message})

}

func (uc *AuthController) RegisterAuthRoutes(rg *gin.RouterGroup) {
	authRoute := rg.Group("/auth")
	
	authRoute.POST("/login", uc.Login)
	authRoute.GET("/check-status",helpers.ValidateJWT,uc.CheckStatus)
}
