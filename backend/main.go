package main

import (
	"github.com/Grvzard/biligank-v2/backend/api"
	"github.com/Grvzard/biligank-v2/backend/crud"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("config.env")
	crud.InitDatabase()
	defer crud.DestoryDatabase()

	r := gin.Default()
	root := r.Group("/")

	api.RegSearch(root)
	api.RegStreamlog(root)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	r.Run(":8080")
}
