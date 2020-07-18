package main

import (
	"context"
	"flag"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Glitchfix/golagobar/config"
	"github.com/Glitchfix/golagobar/db"
	"github.com/Glitchfix/golagobar/server"
	"github.com/Glitchfix/golagobar/services"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var mainContext context.Context
var logWriter io.Writer

func init() {
	time.Local = time.UTC
	logWriter = &lumberjack.Logger{
		Filename: "logs/app.log",
		MaxAge:   30,
		MaxSize:  100,
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, logWriter))
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableLevelTruncation: true,
		TimestampFormat:        time.RFC3339Nano,
	})
	configFile := flag.String("config", "config.json", "Parse the config file")
	mainContext = context.Background()
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logWriter)

	config.Init(configFile)
	db.Init(mainContext)
	services.Init()

}

func main() {
	logrus.Infoln("Start the world")
	server.Init()
}
