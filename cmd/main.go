package main

import (
	_ "github.com/cstungthanh/sandbox/docs"
	"github.com/cstungthanh/sandbox/internal/routes"
	"github.com/cstungthanh/sandbox/pkg/config"
)

// @title Sandbox API DOCUMENTATION
// @version v1.0.0
// @description This is a api for sandbox project.
// @termsOfService http://swagger.io/terms/

// @contact.name Julius
// @contact.url https://github.com/cs-tungthanh
// @contact.email cs.tungthanh@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1

func main() {
	cfg := config.LoadConfig(config.DefaultConfigLoaders())

	r := routes.NewRouter(cfg)

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.Run()
}
