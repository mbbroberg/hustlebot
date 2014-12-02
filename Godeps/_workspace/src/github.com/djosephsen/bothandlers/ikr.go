package bothandlers

import (
	"math/rand"
	"strings"
	"github.com/danryan/hal"
)

type ikr struct{
}

func (h *ikr) Method() string {
	hal.Logger.Debug("ikr: called Method()")
	return hal.HEAR
}

func (h *ikr) Usage() string {
	hal.Logger.Debug("ikr: called Usage()")
	return `ikr - listens for enthusiasm; responds with validation`
}

func (h *ikr) Pattern() string {
	hal.Logger.Debug("ikr: called Pattern()")
	triggers:=[]string{
		"best.*ev(er|ar)",
		"so good",
		"they have the best",
		"awesome",
		"I love",
		"fantastic|wonderful|outstanding|magnificent|brilliant|genius|amazing",
		"ZOMG|OMG|OMFG",
		"(so|pretty) great",
		"off the hook",
	}
	pat := "(?i)"+strings.Join(triggers,"|")
	hal.Logger.Debug("ikr: pattern: %s", pat)
	return pat
}

func (h *ikr) Run(res *hal.Response) error {
	hal.Logger.Debug("ikr: called Run()")
	replies := []string{
		"*I know right?!*",
		"*OMG* couldn't agree more",
		":+1:",
		"+1",
		":arrow_up: THAT",
		":arrow_up: you complete me :arrow_up:",
		"so true",
		"agreed.",
		"that's the fact jack",
		"YUUUUUUP",
		"that's what I'm talkin bout",
		"*IKR?!*",
		"singit",
		"^droppin the truth bombs :boom: :boom: :boom:",
		"#legit",
		"/me nodds emphatically in agreement",
		"for REALZ though",
		"FOR REALSIES",
		"it's like you *literally* just read my mind right now",
	}
	reply := replies[rand.Intn(len(replies)-1)]
	hal.Logger.Debug("ikr: my reply is: %s", reply)
	return res.Send(reply)
}

// Ping exports
var IKR = &ikr{}
