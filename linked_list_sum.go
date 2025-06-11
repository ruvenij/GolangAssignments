package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := l1
	num2 := l2
	broughtForwardVal := 0
	var result *ListNode
	currentNode := result

	for {
		if num1 == nil && num2 == nil && broughtForwardVal == 0 {
			break
		}

		sumVal := broughtForwardVal
		if num1 != nil {
			sumVal += num1.Val
			num1 = num1.Next
		}

		if num2 != nil {
			sumVal += num2.Val
			num2 = num2.Next
		}

		broughtForwardVal = sumVal / 10

		node := &ListNode{
			Val: sumVal % 10,
		}

		if currentNode == nil {
			result = node
			currentNode = node
		} else {
			currentNode.Next = node
			currentNode = currentNode.Next
		}
	}

	return result
}

//func main() {
//	l1 := &ListNode{
//		Val: 2,
//		Next: &ListNode{
//			Val: 4,
//			Next: &ListNode{
//				Val:  3,
//				Next: nil,
//			},
//		},
//	}
//
//	l2 := &ListNode{
//		Val: 5,
//		Next: &ListNode{
//			Val: 6,
//			Next: &ListNode{
//				Val:  4,
//				Next: nil,
//			},
//		},
//	}
//
//	result := addTwoNumbers(l1, l2)
//	for result != nil {
//		fmt.Print(result.Val, " ")
//		result = result.Next
//	}
//
//	fmt.Println()
//
//	l1 = &ListNode{
//		Val: 9,
//		Next: &ListNode{
//			Val: 9,
//			Next: &ListNode{
//				Val: 9,
//				Next: &ListNode{
//					Val: 9,
//					Next: &ListNode{
//						Val: 9,
//						Next: &ListNode{
//							Val: 9,
//							Next: &ListNode{
//								Val:  9,
//								Next: nil,
//							},
//						},
//					},
//				},
//			},
//		},
//	}
//
//	l2 = &ListNode{
//		Val: 9,
//		Next: &ListNode{
//			Val: 9,
//			Next: &ListNode{
//				Val: 9,
//				Next: &ListNode{
//					Val:  9,
//					Next: nil,
//				},
//			},
//		},
//	}
//
//	result = addTwoNumbers(l1, l2)
//	for result != nil {
//		fmt.Print(result.Val, " ")
//		result = result.Next
//	}
//
//	fmt.Println()
//
//	l1 = &ListNode{
//		Val:  0,
//		Next: nil,
//	}
//
//	l2 = &ListNode{
//		Val:  0,
//		Next: nil,
//	}
//
//	result = addTwoNumbers(l1, l2)
//	for result != nil {
//		fmt.Print(result.Val, " ")
//		result = result.Next
//	}
//}
