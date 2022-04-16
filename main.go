package main

import (
	"github.com/gin-gonic/gin"
	"go-jwt-demo/pkg/jwt"
	"go-jwt-demo/pkg/pg"
	"go-jwt-demo/user/repository"
	"go-jwt-demo/user/service"
)

var (
	PGMaster = pg.Config{
		Host: "127.0.0.1",
		User: "postgres",
		Db:   "postgres",
		Pwd:  "postgres",
		Port: "5432",
	}
)

func main() {
	router := gin.Default()
	db := pg.NewPgClient(PGMaster)
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)

	router.POST("login", svc.Login)
	router.Use(jwt.VerifyToken)
	router.GET("info", svc.GetUserInfo)

	router.Run()
}
