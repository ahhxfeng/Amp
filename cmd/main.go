package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ahhxfeng/Amp/configs"
	log "github.com/ahhxfeng/Amp/logger"
)

func Index(c *gin.Context) {
	c.String(200, "Hello world!")
}

func main() {
	// load config file
	// should merge variable declaration with assignment on next line
	// var Conf *configs.Config
	Conf := configs.LoadConfig("../configs/config.yaml")
	fmt.Printf("config: %v\n", *Conf)

	// set logger
	logger := log.InitLogger(Conf.LogConfig.LogFile)
	logger.Info("logger initialized", "filename", Conf.LogConfig.LogFile)
	// simple go http server router use gin
	r := gin.Default()
	r.GET("/index", Index)
	r.GET("/", Index)

	// start the defualt http server 0.0.0.:8080
	r.Run()
}
