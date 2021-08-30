package api

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.POST("/daily/29-08-2021", postD29082021)
	r.POST("/prob/001", post001)
	r.POST("/prob/002", post002)
}
