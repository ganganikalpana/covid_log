package main

import (
	"github.com/ganganikalpana/covidLog/app"
	"github.com/ganganikalpana/covidLog/logger"
)

func main() {
	logger.Info("starting the application")
	app.Start()

}
