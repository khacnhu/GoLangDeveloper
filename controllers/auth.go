package controllers

import (
	"github.com/gin-gonic/gin"
)

type AuthController struct {
}

func (a *AuthController) InitAuthController() *AuthController {
	return a
}

func (a *AuthController) InitAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.GET("/", a.GetAuth())
	router.POST("/login")
	router.POST("/register")
}

func (a *AuthController) GetAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "check auth oke",
		})
	}
}
