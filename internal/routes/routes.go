package routes

import (
	"context"
	"strings"

	"github.com/cstungthanh/sandbox/docs"
	"github.com/cstungthanh/sandbox/pkg/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupCORS(r *gin.Engine, cfg *config.Config) {
	corsOrigins := strings.Split(cfg.ApiServer.AllowedOrigins, ";")
	r.Use(func(c *gin.Context) {
		cors.New(
			cors.Config{
				AllowOrigins: corsOrigins,
				AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
				AllowHeaders: []string{
					"Origin", "Host", "Content-Type", "Content-Length", "Accept-Encoding", "Accept-Language", "Accept",
					"X-CSRF-Token", "Authorization", "X-Requested-With", "X-Access-Token",
				},
				AllowCredentials: true,
			},
		)(c)
	})
}

func NewRouter(ctx context.Context, cfg *config.Config) *gin.Engine {
	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "This is a swagger for API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	r := gin.New()
	pprof.Register(r)

	// r.Use(
	// 	gin.LoggerWithWriter(gin.DefaultWriter, "/healthz"),
	// 	gin.Recovery(),
	// )

	// config CORS
	setupCORS(r, cfg)

	// r.GET("/healthz", h.Healthcheck.Healthz)

	// use ginSwagger middleware to serve the API docs
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// load API here
	// loadV1Routes(r, h, repo, s, cfg)

	return r
}
