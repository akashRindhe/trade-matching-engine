package core

import (
	"log"
	"time"

	"go.uber.org/atomic"
)

var isEngineInitialized = atomic.NewBool(false)
var instruments []Instrument
var isEngineRunning *atomic.Bool
var ordersMatching *atomic.Bool
var orderChannel chan Order
var tradeExecutionChannel chan TradeExecutionReport

/*Initialize - Initializes the engine. Currently the implementation only includes creating orderbooks for each security*/
func Initialize(instruments []Instrument) {
	isEngineInitialized.Store(true)
	isEngineRunning = atomic.NewBool(false)
	ordersMatching = atomic.NewBool(false)
	orderChannel = make(chan Order, 999999)
	tradeExecutionChannel = make(chan TradeExecutionReport, 999999)
}

/*
Start - Starts up the engine allowing for submitting/cancelling/matching of orders.
Allows specification of a time duration to defer matching orders in the case of auction phase
*/
func Start(deferMatchingDuration time.Duration) (orderChannel chan<- Order, tradeExecutionChannel <-chan TradeExecutionReport) {
	if !isEngineInitialized.Load() {
		log.Println("Engine not initialized")
		return
	}
	go acceptOrders()
	go matchOrders(deferMatchingDuration)
	return orderChannel, tradeExecutionChannel
}

func matchOrders(deferMatchingDuration time.Duration) {

}

func acceptOrders() {
	isEngineRunning.Store(true)
	for {
		if !isEngineRunning.Load() {
			break
		}
		//var order = <-orderChannel
		//map[order.SecurityID
	}
}

/*Stop - Stop engine*/
func Stop() {
	if !isEngineRunning.Load() {
		log.Println("Engine not running")
		return
	}
	log.Println("Stopping order matching")
	ordersMatching.Store(false)
	isEngineRunning.Store(false)
}
