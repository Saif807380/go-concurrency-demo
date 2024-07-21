package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	receiveOrders()
	for _, order := range orders {
		fmt.Println(order)
	}
}

var rawOrders = []string{
	`{"productCode": "1111", "quantity": 5, "status": 1}`,
	`{"productCode": "2222", "quantity": 42.3, "status": 1}`,
	`{"productCode": "3333", "quantity": 19, "status": 1}`,
	`{"productCode": "4444", "quantity": 8, "status": 1}`,
}

func receiveOrders() {
	for _, rawOrder := range rawOrders {
		// Decode the JSON string into an order struct
		// using the json.Unmarshal function
		var newOrder order
		err := json.Unmarshal([]byte(rawOrder), &newOrder)
		if err != nil {
			fmt.Println(err)
			continue
		}
		orders = append(orders, newOrder)
	}
}
