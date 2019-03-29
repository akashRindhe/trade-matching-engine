package core

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/atomic"
)

/*Engine - Trade matching engine*/
type Engine struct {
	instrumentChannelMap           map[Instrument]chan<- Order
	instrumentOrderBookMap         map[Instrument]*OrderBook
	isEngineRunning                *atomic.Bool
	ordersMatching                 *atomic.Bool
	internalChannel                chan bool
	orderBookTradeExecutionChannel chan TradeExecutionReport
	clientOrderChannel             chan Order
	clientTradeExecutionChannel    chan TradeExecutionReport
}

/*NewTradeMatchingEngine - Creates a new engine and initializes order books for each security*/
func NewTradeMatchingEngine(instruments []Instrument) Engine {
	e := Engine{
		instrumentChannelMap:           make(map[Instrument]chan<- Order),
		instrumentOrderBookMap:         make(map[Instrument]*OrderBook),
		isEngineRunning:                atomic.NewBool(false),
		ordersMatching:                 atomic.NewBool(false),
		internalChannel:                make(chan bool),
		orderBookTradeExecutionChannel: make(chan TradeExecutionReport, 999999),
		clientOrderChannel:             make(chan Order),
		clientTradeExecutionChannel:    make(chan TradeExecutionReport, 999999)}

	for _, i := range instruments {
		c := make(chan Order, 999999)
		o := NewOrderBook(i, c, e.orderBookTradeExecutionChannel, e.ordersMatching)
		e.instrumentChannelMap[i] = c
		e.instrumentOrderBookMap[i] = o
	}
	return e
}

/*
Start - Starts up the engine allowing for submitting/cancelling/matching of orders.
Allows specification of a time duration to defer matching orders in the case of auction phase.
Returns trade execution channel on which execution reports are returned
*/
func (e Engine) Start(deferMatchingDuration time.Duration) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go e.matchOrders(deferMatchingDuration)
	e.acceptOrders()
	go func() {
		for {
			select {
			case <-sigChan:
				e.internalChannel <- true
			case <-e.internalChannel:
				log.Println("Stopping Engine...")
				return
			}
		}
	}()
}

func (e *Engine) PauseOrderMatching(deferMatchingDuration time.Duration) {

}

/*
Stop - Stop the engine if it is running.
All unmatched orders will be cancelled and trade reports sent back.
*/
func (e *Engine) Stop() {
	e.internalChannel <- true
}

func (e *Engine) matchOrders(deferMatchingDuration time.Duration) {
	log.Println("Sleeping for", deferMatchingDuration.Seconds(), "seconds before starting order matching")
	time.Sleep(deferMatchingDuration)
	log.Println("Starting order matching")
	e.ordersMatching.Store(true)
}

func (e *Engine) acceptOrders() {
	for {
		msg := <-e.clientOrderChannel
		e.instrumentChannelMap[msg.SecurityID] <- msg
	}
}

func (e *Engine) AddOrder(msg Order) {
	e.clientOrderChannel <- msg
}
