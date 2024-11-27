package controller

import (
	"backend/internal/controller/dtos"
	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type AuthController struct {
	authService  service.IAuthService
	tokenService service.ITokenService
}

func ProvideAuthController(authService service.IAuthService, tokenService service.ITokenService) IAuthController {
	return AuthController{
		authService:  authService,
		tokenService: tokenService,
	}
}

func (a AuthController) Login(ctx *gin.Context) {
	var loginRequest dtos.LoginRequest
	err := ctx.ShouldBindBodyWithJSON(&loginRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := a.authService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := a.tokenService.GenerateToken(user.Username)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res := dtos.LoginResponse{
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
	ctx.JSON(200, res)
}

func (a AuthController) Register(ctx *gin.Context) {
	var registerRequest dtos.RegisterRequest
	err := ctx.ShouldBindBodyWithJSON(&registerRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := a.authService.Register(registerRequest.Username, registerRequest.Password, registerRequest.Email)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := a.tokenService.GenerateToken(user.Username)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res := dtos.LoginResponse{
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
	ctx.JSON(200, res)
}
