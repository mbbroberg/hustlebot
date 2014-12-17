package main

import (
	"github.com/djosephsen/hal"
	_ "github.com/djosephsen/hal/adapter/slack"
	"github.com/djosephsen/HalHandlers"
	_ "github.com/djosephsen/hal/store/memory"
	"os"
)

func run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		hal.Logger.Error(err)
		return 1
	}

	robot.Handle(
		HalHandlers.Syn,
		HalHandlers.Tableflip,
		HalHandlers.IKR,
		HalHandlers.ListRooms,
		HalHandlers.ManageChores,
		HalHandlers.ListChores,
		HalHandlers.LoveAndWar,
		HalHandlers.Gifme,
		HalHandlers.Help,
	)

	if err := robot.Run(); err != nil {
		hal.Logger.Error(err)
		return 1
	}
	return 0
}

func main() {
	os.Exit(run())
}
