package main

import (
	"trade-matching-engine/core"
)

func main() {
	ins := []core.Instrument{"AAPL", "IBM"}
	e := core.NewTradeMatchingEngine(ins)
	e.Stop()
}

/*
	A trade matching engine : Software that matches a bid against an ask.

	1) Allows for order to be input (added)
	2) Allows for order to be removed (cancelled)
	3) Matches bids against asks
	4)
*/
