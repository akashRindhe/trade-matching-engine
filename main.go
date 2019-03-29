package main

import (
	"time"
	"trade-matching-engine/core"

	"github.com/robaho/fixed"
)

func main() {
	ins := []core.Instrument{"AAPL", "IBM"}
	e := core.NewTradeMatchingEngine(ins)
	dur, _ := time.ParseDuration("5s")
	e.Start(dur)
	time.Sleep(25 * time.Second)
	msg := core.Order{
		SecurityID: "AAPL",
		OrderID:    "ABC123",
		Px:         fixed.NewF(10.5),
		Qty:        fixed.NewI(100, 0),
		Side:       core.BUY,
		EntryTime:  time.Now()}
	e.AddOrder(msg)

	//go e.PauseOrderMatching(dur)
	//time.Sleep(60 * time.Second)
	//e.Stop()
	//defer e.Stop()
}

/*
	A trade matching engine : Software that matches a bid against an ask.

	1) Allows for order to be input (added)
	2) Allows for order to be removed (cancelled)
	3) Matches bids against asks
	4)
*/
