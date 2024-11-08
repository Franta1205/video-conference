package routes

import (
	"github.com/gin-gonic/gin"
	"video-conference/controllers"
)

func SetUpRoutes(router *gin.Engine) {
	homeController := controllers.NewHomeController()
	router.GET("/", homeController.Index)

	websocketController := controllers.NewWebsocketsController()
	router.GET("/ws", websocketController.HandleWebsocket)

	callController := controllers.NewCallController()
	callGroup := router.Group("/call")
	{
		callGroup.POST("/start", callController.StartCall)
		callGroup.POST("/join", callController.JoinCall)
		callGroup.POST("/end", callController.EndCall)
		callGroup.GET("/page", callController.CallPage)
	}
}
