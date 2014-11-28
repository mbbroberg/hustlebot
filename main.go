package main

import (
	"github.com/danryan/hal"
	_ "github.com/danryan/hal/adapter/slack"
	"github.com/djosephsen/bothandlers"
	_ "github.com/danryan/hal/store/memory"
	"os"
)

func run() int {
	robot, err := hal.NewRobot()
	if err != nil {
		hal.Logger.Error(err)
		return 1
	}

	robot.Handle(
		bothandlers.Syn,
		bothandlers.Tableflip,
		bothandlers.IKR,
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
