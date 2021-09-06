package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	s002 "github.com/viardant/leetcode-solutions/0002"
)

type post002REQUEST struct {
	L1 s002.ListNode `json:"l1" binding:"required"`
	L2 s002.ListNode `json:"l2" binding:"required"`
}

type post002RESPONSE struct {
	R s002.ListNode `json:"res-node"`
}

func post002(c *gin.Context) {
	log := c.MustGet("log").(*logrus.Logger)

	req := post002REQUEST{}
	if err := c.Bind(&req); err != nil {
		log.WithError(err).Error("could not bind request")
		return
	}

	c.JSON(http.StatusOK, post002RESPONSE{R: *s002.AddTwoNumbers(&req.L1, &req.L2)})
}
