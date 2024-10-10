package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ExquisiteCore/LagrangeGo-Template/bot"
	"github.com/ExquisiteCore/LagrangeGo-Template/config"
	"github.com/ExquisiteCore/LagrangeGo-Template/logic"
	"github.com/ExquisiteCore/LagrangeGo-Template/utils"
)

func init() {
	config.Init()
	utils.Init()
	bot.Init()
}

func main() {

	bot.Login()

	logic.RegisterCustomLogic()

	logic.SetupLogic()

	defer bot.QQClient.Release()

	defer bot.Dumpsig()

	// setup the main stop channel

	mc := make(chan os.Signal, 2)
	signal.Notify(mc, os.Interrupt, syscall.SIGTERM)
	for {
		switch <-mc {
		case os.Interrupt, syscall.SIGTERM:
			return
		}
	}
}
