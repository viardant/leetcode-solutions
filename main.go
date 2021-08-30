package main

import (
	"github.com/sirupsen/logrus"
	"github.com/viardant/leetcode-solutions/webapi/api"
)

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	router := api.Handler()

}
