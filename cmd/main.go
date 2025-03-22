package main

import (
	"github.com/gin-gonic/gin"
	"gravitum-test/internal/user"
	"gravitum-test/pkg"
)

func init() {
	pkg.InitDB()
}

func main() {
	repo := user.NewRepository()
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	r := gin.Default()
	r.POST("/users", handler.CreateUser)
	r.GET("/users/:id", handler.GetUser)
	r.PUT("/users/:id", handler.UpdateUser)

	r.Run(":8080")
}
