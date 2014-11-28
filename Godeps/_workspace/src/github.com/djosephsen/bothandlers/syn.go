package bothandlers

import (
	"github.com/danryan/hal"
)

var Syn = hal.Respond(`syn`, func(res *hal.Response) error {
	return res.Reply(`ack`)
})
