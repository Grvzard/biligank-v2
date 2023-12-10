package main

import (
	"log"
	"os"
	"strings"

	"github.com/Grvzard/biligank-v2/backend/api"
	"github.com/Grvzard/biligank-v2/backend/crud"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("config.env")

	crud.InitDatabase()
	defer crud.DestoryDatabase()

	r := gin.Default()

	cors_s := strings.Split(os.Getenv("API_CORS"), ",")
	log.Print("CORS: ", cors_s)

	r.Use(cors.New(cors.Config{
		AllowOrigins: cors_s,
	}))

	root := r.Group("/")
	api.RegSearch(root)
	api.RegStreamlog(root)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	r.Run(":8080")
}
