package core

import (
	"go.uber.org/atomic"
)

type OrderBook struct {
	securityID            Instrument
	orderChannel          <-chan Order
	tradeExecutionChannel chan<- TradeExecutionReport
	matchOrders           *atomic.Bool
}

func NewOrderBook(securityID Instrument,
	orderChannel <-chan Order,
	tradeExecutionChannel chan<- TradeExecutionReport,
	matchOrders *atomic.Bool) *OrderBook {
	o := OrderBook{securityID: securityID,
		orderChannel:          orderChannel,
		tradeExecutionChannel: tradeExecutionChannel,
		matchOrders:           matchOrders}
	return &o
}

func (o *OrderBook) Start() {

}
