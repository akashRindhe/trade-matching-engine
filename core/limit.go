package core

import (
	"github.com/robaho/fixed"
)

type Limit struct {
	LmtPx fixed.Fixed
	Qty   fixed.Fixed
	Head  *Order
}
