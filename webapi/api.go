package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Handler(log *logrus.Logger) http.Handler {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("log", log)
		c.Next()
	})

	return router
}
