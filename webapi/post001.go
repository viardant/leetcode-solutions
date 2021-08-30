package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	s001 "github.com/viardant/leetcode-solutions/001"
)

type post001REQUEST struct {
	Nums   []int `json:"nums" binding:"required"`
	Target int   `json:"target" binding:"required"`
}

type post001RESPONSE struct {
	Res []int `json:"positions"`
}

func post001(c *gin.Context) {
	log := c.MustGet("log").(*logrus.Logger)

	req := post001REQUEST{}
	if err := c.Bind(&req); err != nil {
		log.WithError(err).Error("could not bind request")
		return
	}

	c.JSON(http.StatusOK, post001RESPONSE{Res: s001.TwoSum(req.Nums, req.Target)})
}
