package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	config "github.com/maykemrezende/movies-recommendation-poc/config"
	docs "github.com/maykemrezende/movies-recommendation-poc/docs"
	"github.com/maykemrezende/movies-recommendation-poc/internal/movies/handlers"
	"github.com/maykemrezende/movies-recommendation-poc/internal/movies/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	application, err := config.LoadApplication()

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	r := InitRoutes(application.Db, application.RootUrl)

	r.Run(fmt.Sprintf(":%s", application.ApiPort))
}

func InitRoutes(db *gorm.DB, rootUrl string) *gin.Engine {
	r := gin.Default()
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Couldn't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	// TODO: remove CORS allows all origins
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	r.Use(cors.New(conf))

	docs.SwaggerInfo.BasePath = "/api/v1"

	routes.InitMoviesRoutes(r, handlers.NewMoviesHandler(logger))

	formatedUrl := fmt.Sprintf("%s/swagger/doc.json", rootUrl)
	url := ginSwagger.URL(formatedUrl)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	return r
}
