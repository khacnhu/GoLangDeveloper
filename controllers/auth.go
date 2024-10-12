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
	auth := router.Group("/auth")
	auth.GET("/", a.GetAuth())
	router.POST("/login")
	auth.POST("/register", a.AuthRegister())
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
		Id       int    `json:"id"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:required"`
	}

	return func(ctx *gin.Context) {
		var registerBody RegisterBody

		if err := ctx.BindJSON(&registerBody); err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}

		user, err := a.authService.Register(&registerBody.Id, &registerBody.Email, &registerBody.Password)

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
