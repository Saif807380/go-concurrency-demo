package main

import "fmt"

type order struct {
	ProductCode string
	Quantity    float64
	Status      orderStatus
}

func (o order) String() string {
	return fmt.Sprintf("Product Code: %s, Quantity: %.2f, Status: %v", o.ProductCode, o.Quantity, orderStatusToText(o.Status))
}

func orderStatusToText(status orderStatus) string {
	switch status {
	case none:
		return "None"
	case new:
		return "New"
	case received:
		return "Received"
	case reserved:
		return "Reserved"
	case filled:
		return "Filled"
	default:
		return "Unknown status"
	}
}

type orderStatus int

const (
	none orderStatus = iota
	new
	received
	reserved
	filled
)

var orders []order
