package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	d29082021 "github.com/viardant/leetcode-solutions/D-29-08-2021"
)

type postD29082021REQUEST struct {
	Nums []int `json:"nums" binding:"required"`
	N    int   `json:"n" binding:"required"`
}

type postD29082021RESPONSE struct {
	N int `json:"patches-required"`
}

func postD29082021(c *gin.Context) {
	log := c.MustGet("log").(*logrus.Logger)

	req := postD29082021REQUEST{}
	if err := c.Bind(&req); err != nil {
		log.WithError(err).Error("could not bind request")
		return
	}

	c.JSON(http.StatusOK, postD29082021RESPONSE{N: d29082021.MinPatches(req.Nums, req.N)})
}
