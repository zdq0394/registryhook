package hook

import "github.com/gin-gonic/gin"

func setRoutes(r *gin.Engine) {
	r.GET("/ping", pingHanlder)
	r.POST("/events", eventsHandler)
}
