package server

import (
	"github.com/Glitchfix/golagobar/handler"
	"github.com/Glitchfix/golagobar/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func chainHandler(h gin.HandlerFunc) gin.HandlersChain {
	return gin.HandlersChain{h, middlewares.ErrorMiddleware}
}

// NewRouter - Create new router object
func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.Default())

	router.GET("/ping", chainHandler(handler.Ping)...)

	openGroup := router.Group("/o")
	{
		openGroup.POST("/login", chainHandler(handler.LoginHandler)...)
		openGroup.POST("/register", chainHandler(handler.RegisterHandler)...)
	}
	restrictedGroup := router.Group("/r")
	{
		restrictedGroup.Use(middlewares.AuthMiddleware)

		users := restrictedGroup.Group("/users")
		{
			users.GET("/profile", chainHandler(handler.ProfileHandler)...)
			users.PUT("/location", chainHandler(handler.SetLocation)...)
			users.POST("/editpass", chainHandler(handler.EditPassword)...)
		}

		rides := restrictedGroup.Group("/rides")
		{
			rides.POST("/estimate", chainHandler(handler.RdieEstimateHandler)...)
			rides.POST("/nearby", chainHandler(handler.NearbyRidesHandler)...)
			rides.PUT("/create", chainHandler(handler.CreateRideHandler)...)
			rides.POST("/accept", chainHandler(handler.AcceptRideHandler)...)
			rides.POST("/cancel", chainHandler(handler.CancelRideHandler)...)
			rides.POST("/complete", chainHandler(handler.CompleteRideHandler)...)
		}
	}
	middlewares.Init()

	return router
}
