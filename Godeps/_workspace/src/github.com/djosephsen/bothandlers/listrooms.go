package bothandlers

import (
	"github.com/djosephsen/hal"
	"fmt"
	)

var Listroom = &hal.Handler{
	Method:  hal.RESPOND,
	Pattern: `what room is this`,
	Run: func(res *hal.Response) error {
		room := res.Message.Room
		reply := fmt.Sprintf("Current room is: %s",room)
		return res.Send(reply)
	},
}
