package main

import (
	"fmt"
	rbt "github.com/emirpasic/gods/trees/redblacktree"
)

type OrderBookWithRbt struct {
	sellOrders *SellOrders
	buyOrders  *BuyOrders
}

type SellOrders struct {
	orders *rbt.Tree
}

type BuyOrders struct {
	orders *rbt.Tree
}

func NewOrderBookWithRbt() *OrderBookWithRbt {
	return &OrderBookWithRbt{
		sellOrders: &SellOrders{
			orders: rbt.NewWith(floatComparator),
		},
		buyOrders: &BuyOrders{
			orders: rbt.NewWith(floatComparator),
		},
	}
}

func floatComparator(a, b interface{}) int {
	da := a.(float64)
	db := b.(float64)

	switch {
	case da < db:
		return -1
	case da > db:
		return 1
	default:
		return 0
	}
}

func (ob *OrderBookWithRbt) InsertOrder(inputOrder *Order) {
	ordType := "SELL"
	if inputOrder.Type == BUY {
		ordType = "BUY"
	}
	fmt.Printf("Input %s Order : %d @ %f\n", ordType, inputOrder.Quantity, inputOrder.Price)
	if inputOrder.Type == BUY {
		book := ob.sellOrders.orders
		keys := book.Keys()
		matchedQuantity := 0
		if len(keys) > 0 {
			for {
				leftNode := book.Left()
				if leftNode == nil || leftNode.Key.(float64) > inputOrder.Price {
					fmt.Println("No sell orders to match from sell orderbook")
					break
				}

				existingOrdersOfLeftNode := leftNode.Value.([]*Order)
				for len(existingOrdersOfLeftNode) > 0 {
					currentMatchedOrder := existingOrdersOfLeftNode[0]

					// possible match
					currentMatchQuantity := MinVal(currentMatchedOrder.Quantity, inputOrder.Quantity)
					matchedQuantity += currentMatchQuantity
					fmt.Printf("Matched with sell order %d @ %f\n", currentMatchQuantity, leftNode.Key)

					// update existing order
					currentMatchedOrder.Quantity -= currentMatchQuantity
					if currentMatchedOrder.Quantity == 0 {
						// remove order from orderbook
						existingOrdersOfLeftNode = existingOrdersOfLeftNode[1:]
						fmt.Printf("Removing sell order %d @ %f\n", currentMatchedOrder.Quantity,
							currentMatchedOrder.Price)
					}

					inputOrder.Quantity -= currentMatchQuantity
					if inputOrder.Quantity == 0 {
						if len(existingOrdersOfLeftNode) == 0 {
							book.Remove(leftNode)
						}

						fmt.Printf("Fully filled buy order %d @ %f\n", inputOrder.Quantity,
							inputOrder.Price)
						return
					}
				}

				if len(existingOrdersOfLeftNode) == 0 {
					fmt.Printf("Removing sell key from buy orderbook %d\n", leftNode.Key)
					book.Remove(leftNode.Key)
				}
			}
		}

		if inputOrder.Quantity > 0 {
			// add to orderbook
			fmt.Printf("Adding buy order to buy orderbook %d @ %f\n", inputOrder.Quantity, inputOrder.Price)
			buyBook := ob.buyOrders.orders
			buyBook.Put(inputOrder.Price, []*Order{inputOrder})
		}
	} else if inputOrder.Type == SELL {
		book := ob.buyOrders.orders
		keys := book.Keys()
		matchedQuantity := 0
		if len(keys) > 0 {
			for {
				rightNode := book.Right()
				if rightNode == nil || rightNode.Key.(float64) < inputOrder.Price {
					fmt.Println("No buy orders to match from buy orderbook")
					break
				}

				existingOrdersOfLeftNode := rightNode.Value.([]*Order)
				for len(existingOrdersOfLeftNode) > 0 {
					currentMatchedOrder := existingOrdersOfLeftNode[0]

					// possible match
					currentMatchQuantity := MinVal(currentMatchedOrder.Quantity, inputOrder.Quantity)
					matchedQuantity += currentMatchQuantity
					fmt.Printf("Matched with buy order %d @ %f\n", currentMatchQuantity, rightNode.Key)

					// update existing order
					currentMatchedOrder.Quantity -= currentMatchQuantity
					if currentMatchedOrder.Quantity == 0 {
						// remove order from orderbook
						existingOrdersOfLeftNode = existingOrdersOfLeftNode[1:]
						fmt.Printf("Removing buy order %d @ %f\n", currentMatchedOrder.Quantity,
							currentMatchedOrder.Price)
					}

					inputOrder.Quantity -= currentMatchQuantity
					if inputOrder.Quantity == 0 {
						if len(existingOrdersOfLeftNode) == 0 {
							book.Remove(rightNode)
						}
						fmt.Printf("Fully filled sell order %d @ %f\n", inputOrder.Quantity,
							inputOrder.Price)

						return
					}
				}

				if len(existingOrdersOfLeftNode) == 0 {
					fmt.Printf("Removing buy key from sell orderbook %d\n", rightNode.Key)
					book.Remove(rightNode)
				}
			}
		}

		if inputOrder.Quantity > 0 {
			// add to orderbook
			fmt.Printf("Adding order to sell orderbook %d @ %f\n", inputOrder.Quantity,
				inputOrder.Price)
			sellBook := ob.sellOrders.orders
			sellBook.Put(inputOrder.Price, []*Order{inputOrder})
		}
	}
}
