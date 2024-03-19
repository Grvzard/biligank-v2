package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/Grvzard/biligank-v2/backend/api"
	"github.com/Grvzard/biligank-v2/backend/config"
	"github.com/Grvzard/biligank-v2/backend/crud"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	godotenv.Load("config.env")

	st1, err := time.Parse("2006_01", os.Getenv("API_STAGE1"))
	if err != nil {
		panic(err)
	}
	var beijing = time.FixedZone("UTC+8", 8*60*60)
	config.GlobalConfig.Stage1 = st1.In(beijing)

	crud.InitDatabase()
	defer crud.DestoryDatabase()

	r := gin.Default()

	cors_s := strings.Split(os.Getenv("API_CORS"), ",")
	log.Print("CORS: ", cors_s)

	r.Use(cors.New(cors.Options{
		AllowedOrigins: cors_s,
	}))

	root := r.Group("/")
	api.RegSearch(root)
	api.RegStreamlog(root)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	r.Run(":8080")
}
