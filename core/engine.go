package core

import (
	"log"
	"time"

	"go.uber.org/atomic"
)

type Engine struct {
	instruments           []Instrument
	isEngineRunning       *atomic.Bool
	ordersMatching        *atomic.Bool
	orderChannel          chan Order
	tradeExecutionChannel chan TradeExecutionReport
}

/*NewTradeMatchingEngine - Creates a new engine and initializes order books for each security*/
func NewTradeMatchingEngine(instruments []Instrument) Engine {
	e := Engine{instruments: instruments,
		isEngineRunning:       atomic.NewBool(false),
		ordersMatching:        atomic.NewBool(false),
		orderChannel:          make(chan Order, 999999),
		tradeExecutionChannel: make(chan TradeExecutionReport, 999999)}
	return e
}

/*
Start - Starts up the engine allowing for submitting/cancelling/matching of orders.
Allows specification of a time duration to defer matching orders in the case of auction phase
*/
func (e *Engine) Start(deferMatchingDuration time.Duration) (orderChannel chan<- Order, tradeExecutionChannel <-chan TradeExecutionReport) {
	go e.acceptOrders()
	go e.matchOrders(deferMatchingDuration)
	return e.orderChannel, e.tradeExecutionChannel
}

/*
Stop - Stop the engine if it is running.
All unmatched orders will be cancelled and trade reports sent back.
*/
func (e *Engine) Stop() {
	if !e.isEngineRunning.Load() {
		log.Println("Engine not running")
		return
	}
	log.Println("Stopping order matching")
	e.ordersMatching.Store(false)
	e.isEngineRunning.Store(false)
}

func (e *Engine) matchOrders(deferMatchingDuration time.Duration) {
	log.Println("Sleeping for ", deferMatchingDuration.Seconds())
	time.Sleep(deferMatchingDuration * time.Second)
	log.Println("Starting order matching")
	e.ordersMatching.Store(true)
}

func (e *Engine) acceptOrders() {
	e.isEngineRunning.Store(true)
	for {
		if !e.isEngineRunning.Load() {
			break
		}
		//var order = <-orderChannel
		//map[order.SecurityID
	}
}
