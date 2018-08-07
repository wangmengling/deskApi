package main

import (
	"deskapi/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	//"time"
	"deskapi/config"
)

func init() {
	db.Connect()
}

func main() {
	router := gin.Default()
	router.Static("static", "./static")
	router.Use(Connect)
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"PUT", "PATCH","POST","GET"},
	//	AllowHeaders:     []string{"Origin"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	MaxAge: 12 * time.Hour,
	//}))
	router.Use(cors.Default())
	router.Use(ErrorHandler)
	routes.Routes(router)

	router.Run(":3002") // listen and serve on 0.0.0.0:8080
}

func Connect(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}

// ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
	c.Next()

	// TODO: Handle it in a better way
	if len(c.Errors) > 0 {
		c.HTML(http.StatusBadRequest, "400", gin.H{
			"errors": c.Errors,
		})
	}
}