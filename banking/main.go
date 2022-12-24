package main

import (
	"banking/app"
	"banking/logger"
)

func main() {
	//log.Println("Starting app...")
	logger.Info("Starting app...")
	app.Start()
}
