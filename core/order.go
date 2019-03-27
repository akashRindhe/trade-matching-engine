package core

import (
	"time"

	"github.com/robaho/fixed"
)

type Order struct {
	OrderID    string
	Px         fixed.Fixed
	Qty        fixed.Fixed
	SecurityID Instrument
	EntryTime  time.Time
}
