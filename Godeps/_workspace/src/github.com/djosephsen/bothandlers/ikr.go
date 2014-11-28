package bothandlers

import (
	"math/rand"
	"strings"
	"github.com/danryan/hal"
)

type ikr struct{
}

func (h *ikr) Method() string {
	return hal.RESPOND
}

func (h *ikr) Usage() string {
	return `ping - responds with "PONG"`
}

func (h *ikr) Pattern() string {
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
	return "(?i)"+strings.Join(triggers,"|")
}

func (h *ikr) Run(res *hal.Response) error {
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

	return res.Send(replies[rand.Intn(len(replies))])
}

// Ping exports
var IKR = &ikr{}
