package core

import (
	"time"

	"github.com/robaho/fixed"
)

type Side int

const (
	BUY  Side = iota + 1
	SELL Side = iota
)

type Order struct {
	OrderID    string
	Px         fixed.Fixed
	Qty        fixed.Fixed
	SecurityID Instrument
	EntryTime  time.Time
	Side       Side
}
