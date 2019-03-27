package core

import (
	"time"

	"github.com/robaho/fixed"
)

type TradeExecutionReport struct {
	OrderID    string
	MatchID    string
	Px         fixed.Fixed
	Qty        fixed.Fixed
	SecurityID Instrument
	MatchTime  time.Time
	EntryTime  time.Time
}
