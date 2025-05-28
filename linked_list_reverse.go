package main

import "fmt"

type LinkedNode struct {
	Value int
	Next  *LinkedNode
}

func RevertLinkedList(currentHead *LinkedNode) *LinkedNode {
	currentNode := currentHead
	var prev *LinkedNode

	for currentNode != nil {
		fmt.Println("Current : ", currentNode.Value)
		next := currentNode.Next
		currentNode.Next = prev

		prev = currentNode
		currentNode = next
	}

	return prev
}

func IterateLinkedList(currentHead *LinkedNode) {
	currentNode := currentHead
	for currentNode.Next != nil {
		fmt.Println("Current Node : ", currentNode.Value)
		currentNode = currentNode.Next
	}

	fmt.Println("Current Node : ", currentNode.Value)
}

// IsCyclePresent using floyd's cycle detection algorithm
func IsCyclePresent(head *LinkedNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			fmt.Println("A cycle is present")
			return true
		}
	}

	fmt.Println("A cycle is not present")
	return false
}
