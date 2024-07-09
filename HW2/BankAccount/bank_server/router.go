package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/account/:id", GetAccount)
	r.POST("/account", CreateAccount)
	r.PUT("/account/:id/name", UpdateAccountName)
	r.PUT("/account/:id/balance", UpdateAccountBalance)
	r.DELETE("/account/:id", DeleteAccount)

	return r
}
