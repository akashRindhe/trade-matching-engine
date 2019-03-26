package core

import "log"

var instruments []Instrument
var isEngineStarted bool
var isOrderMatchingStarted bool

/*Initialize - Initializes the engine. Currently the implementation only includes creating orderbooks for each security*/
func Initialize(instruments []Instrument) {

}

/*Start - Starts up the engine allowing for submitting/cancelling/matching of orders. If matching is to be deferred, use StartWithoutMatching*/
func Start() {
	isEngineStarted = true
	isOrderMatchingStarted = true
}

/*StartWithoutOrderMatching - Starts up the engine allowing for submitting/cancelling of orders.*/
func StartWithoutOrderMatching() {
	isEngineStarted = true
}

/*StartOrderMatching - Start order matching. The engine has to be in the started state else error is thrown*/
func StartOrderMatching() {
	if !isEngineStarted {
		log.Fatal("Trade matching engine is down")
	}
	isOrderMatchingStarted = true
	go matchOrders()
}

func matchOrders() {

}

func SubmitOrder(order Order) {
	if !isEngineStarted {
		log.Fatal("Trade matching engine is down")
	}
}

func CancelOrder(orderId string) {
	if !isEngineStarted {
		log.Fatal("Trade matching engine is down")
	}
}
