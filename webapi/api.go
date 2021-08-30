package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Handler(log *logrus.Logger) *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("log", log)
		c.Next()
	})

	return router
}
