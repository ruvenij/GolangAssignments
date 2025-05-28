package main

import (
	"container/heap"
	"fmt"
)

// OrderType defines BUY or SELL
type OrderType int

const (
	BUY OrderType = iota
	SELL
)

// Order represents a single order
type Order struct {
	ID       int
	Price    float64
	Quantity int
	Type     OrderType
}

// AskHeap implements a Min-Heap for sell orders (asks)
type AskHeap []*Order

func (h AskHeap) Len() int           { return len(h) }
func (h AskHeap) Less(i, j int) bool { return h[i].Price < h[j].Price }
func (h AskHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *AskHeap) Push(x interface{}) {
	*h = append(*h, x.(*Order))
}

func (h *AskHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// BidHeap implements a Max-Heap for buy orders (bids)
type BidHeap []*Order

func (h BidHeap) Len() int           { return len(h) }
func (h BidHeap) Less(i, j int) bool { return h[i].Price > h[j].Price }
func (h BidHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *BidHeap) Push(x interface{}) {
	*h = append(*h, x.(*Order))
}

func (h *BidHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// OrderBook contains bids and asks
type OrderBook struct {
	Bids BidHeap
	Asks AskHeap
}

func NewOrderBook() *OrderBook {
	b := &BidHeap{}
	a := &AskHeap{}
	heap.Init(b)
	heap.Init(a)
	return &OrderBook{Bids: *b, Asks: *a}
}

// AddOrder processes and matches an order
func (ob *OrderBook) AddOrder(order *Order) {
	switch order.Type {
	case BUY:
		// Try to match with lowest ask
		for order.Quantity > 0 && ob.Asks.Len() > 0 {
			bestAsk := ob.Asks[0]
			if bestAsk.Price > order.Price {
				break // No match
			}

			// Match
			matchQty := MinVal(order.Quantity, bestAsk.Quantity)
			fmt.Printf("Matched BUY %d @ %.2f with SELL %d @ %.2f\n",
				matchQty, bestAsk.Price, matchQty, bestAsk.Price)

			order.Quantity -= matchQty
			bestAsk.Quantity -= matchQty

			if bestAsk.Quantity == 0 {
				heap.Pop(&ob.Asks)
			}
		}
		if order.Quantity > 0 {
			heap.Push(&ob.Bids, order)
		}

	case SELL:
		// Try to match with highest bid
		for order.Quantity > 0 && ob.Bids.Len() > 0 {
			bestBid := ob.Bids[0]
			if bestBid.Price < order.Price {
				break // No match
			}

			// Match
			matchQty := MinVal(order.Quantity, bestBid.Quantity)
			fmt.Printf("Matched SELL %d @ %.2f with BUY %d @ %.2f\n",
				matchQty, bestBid.Price, matchQty, bestBid.Price)

			order.Quantity -= matchQty
			bestBid.Quantity -= matchQty

			if bestBid.Quantity == 0 {
				heap.Pop(&ob.Bids)
			}
		}
		if order.Quantity > 0 {
			heap.Push(&ob.Asks, order)
		}
	}
}

func MinVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
