package main


import (
	"duidQu/db"
    "duidQu/handler/api"
    "duidQu/repository"
    "duidQu/service"

    "github.com/gin-gonic/gin"
)


func main() {
	//init db
	db.Init()
	
	//init repository, service, handler
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)


	//setup router
	r := gin.Default()
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserByID)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	//jalankan server
	r.Run(":8080")


}