package controllers

import (
	"go-tutorial/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func (a *AuthController) InitAuthController(authService services.AuthService) *AuthController {
	a.authService = authService
	return a
}

func (a *AuthController) InitAuthRoutes(router *gin.Engine) {
	authRouter := router.Group("/auth")
	authRouter.GET("/", a.GetAuth())
	authRouter.POST("/login", a.AuthLogin())
	authRouter.POST("/register", a.AuthRegister())
}

func (a *AuthController) GetAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "check auth oke",
		})
	}
}

func (a *AuthController) AuthRegister() gin.HandlerFunc {

	type RegisterBody struct {
		Email    string `json:"email" gorm:"unique;not null" binding:"required"`
		Password string `json:"password"`
	}

	return func(ctx *gin.Context) {
		var registerBody RegisterBody

		if err := ctx.ShouldBindJSON(&registerBody); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := a.authService.Register(&registerBody.Email, &registerBody.Password)

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"user": user,
		})

	}
}

func (a *AuthController) AuthLogin() gin.HandlerFunc {

	type LoginBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password"`
	}

	return func(ctx *gin.Context) {
		var loginBody LoginBody

		if err := ctx.ShouldBindJSON(&loginBody); err != nil {
			ctx.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := a.authService.Login(&loginBody.Email, &loginBody.Password)

		if err != nil {
			ctx.JSON(404, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"user": user,
		})

	}
}
