package core

import (
	"time"

	. "github.com/robaho/fixed"
)

type TradeExecutionReport struct {
	OrderID    string
	MatchID    string
	Px         Fixed
	Qty        Fixed
	SecurityID Instrument
	MatchTime  time.Time
	EntryTime  time.Time
}
