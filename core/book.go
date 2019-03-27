package core

import "sync"

type OrderBook struct {
	securityID            Instrument
	orderChannel          <-chan Order
	tradeExecutionChannel chan<- TradeExecutionReport
	orderMatchingLock     *sync.Cond
}

func NewOrderBook(securityID Instrument,
	orderChannel <-chan Order,
	tradeExecutionChannel chan<- TradeExecutionReport,
	orderMatchingLock *sync.Cond) OrderBook {
	o := OrderBook{securityID: securityID,
		orderChannel:          orderChannel,
		tradeExecutionChannel: tradeExecutionChannel,
		orderMatchingLock:     orderMatchingLock}
	return o
}
