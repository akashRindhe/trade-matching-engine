package core

import (
	. "github.com/robaho/fixed"
)

type Limit struct {
	LmtPx Fixed
	Qty   Fixed
	Head  *Order
}
