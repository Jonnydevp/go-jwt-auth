package main

import (
	"go-jwt-auth/controllers"
	"go-jwt-auth/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("signup", controllers.Signup)
	r.Run()
}
