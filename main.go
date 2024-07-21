package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

/*
receiveOrders function reads rawOrders and sends each order to a channel
The channel is read by validateOrders which does some validation
Valid orders are sent to validOrderCh and invalid orders are sent to invalidOrderCh
Finally, the main goroutine reads from both channels and prints the orders
*/
func main() {
	var wg sync.WaitGroup

	// initialize channels
	var receivedOrdersCh = make(chan order)
	var validOrderCh = make(chan order)
	var invalidOrderCh = make(chan invalidOrder)

	// start goroutines
	go receiveOrders(receivedOrdersCh)
	go validateOrders(receivedOrdersCh, validOrderCh, invalidOrderCh)

	wg.Add(1)
	go func(validOrderCh <-chan order, invalidOrderCh <-chan invalidOrder) {
	loop:
		for {
			select {
			case order, ok := <-validOrderCh:
				if ok {
					fmt.Printf("Valid Order received: %v\n", order)
					orders = append(orders, order)
				} else {
					break loop
				}
			case invalidOrder, ok := <-invalidOrderCh:
				if ok {
					fmt.Printf("Invalid Order received: %v\n", invalidOrder)
				} else {
					break loop
				}
			}

		}
		wg.Done()
	}(validOrderCh, invalidOrderCh)

	wg.Wait()
	fmt.Println("Printing valid received orders")
	for _, order := range orders {
		fmt.Println(order)
	}
}

var rawOrders = []string{
	`{"productCode": "1111", "quantity": 5, "status": 1}`,
	`{"productCode": "2222", "quantity": 42.3, "status": 1}`,
	`{"productCode": "3333", "quantity": -19, "status": 1}`,
	`{"productCode": "4444", "quantity": 8, "status": 1}`,
}

func validateOrders(in <-chan order, out chan<- order, errCh chan<- invalidOrder) {
	for order := range in {
		if order.Quantity < 0 {
			errCh <- invalidOrder{order, fmt.Errorf("quantity is less than 0")}
		} else {
			out <- order
		}
	}
	close(out)
	close(errCh)
}

func receiveOrders(out chan<- order) {
	for _, rawOrder := range rawOrders {
		// Decode the JSON string into an order struct
		// using the json.Unmarshal function
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			fmt.Println(err)
			continue
		}
		out <- newOrder
	}
	close(out)
}
