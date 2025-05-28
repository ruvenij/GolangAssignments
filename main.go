package main

import (
	"fmt"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func producer(ch chan string) {
	time.Sleep(1 * time.Second)
	fmt.Println("Producer passing")
	ch <- "Hello from producer"
	fmt.Println("Producer done")
}

func consumer(ch chan string) {
	msg := <-ch
	fmt.Println("Received:", msg)
}

func main() {
	//s := "gopher"
	//fmt.Printf("Hello and welcome, %s!\n", s)
	//
	//for i := 1; i <= 5; i++ {
	//fmt.Println("i =", 100/i)
	//}

	//fmt.Println(LongestCommonPrefix([]string{"abc", "def", "ghi"}))
	//arrayInput := []int{0, 1, 2, 2, 3, 0, 4, 2}
	////fmt.Println(RemoveDuplicates(arrayInput))
	////fmt.Println(arrayInput)
	//
	//arrayInput = []int{0, 1, 2, 2, 3, 0, 4, 2}
	//fmt.Println(arrayInput)
	//fmt.Println(RemoveElement(arrayInput, 2))
	//fmt.Println(arrayInput)
	//
	//arrayInput = []int{4, 2, 0, 2, 2, 1, 4, 4, 1, 4, 3, 2}
	//fmt.Println(arrayInput)
	//fmt.Println(RemoveElement(arrayInput, 4))
	//fmt.Println(arrayInput)
	//
	//fmt.Println(StrStr("sadbutsad", "sad"))
	//fmt.Println(StrStr("leetcode", "leeto"))
	//fmt.Println(StrStr("a", "a"))
	//
	//mySlice := make([]int, 5)
	//fmt.Println(mySlice)
	//mySlice = append(mySlice, 1)
	//fmt.Println(mySlice)
	//
	////ch := make(chan string)
	////go producer(ch)
	////go consumer(ch)
	////
	////time.Sleep(2 * time.Second) // wait for goroutines to finish
	//
	//// go routine
	//chan1 := make(chan string)
	//go PublishMessages(chan1)
	//go ReadMessages(chan1)
	//time.Sleep(2 * time.Second)
	//
	//// sum
	//sum := FindSum(1, 100)
	//fmt.Println("Sum of even numbers : ", sum)
	//
	//// min max
	//elem := []int{5, 8, 2, 10, 3}
	//minElem, maxElem, _ := FindMinMaxWithSort(elem)
	//fmt.Println("Minimum element : ", minElem)
	//fmt.Println("Maximum element : ", maxElem)
	//
	//minElem, maxElem, _ = FindMinMaxWithRange(elem)
	//fmt.Println("Minimum element : ", minElem)
	//fmt.Println("Maximum element : ", maxElem)
	//
	//// prime
	//isPrime := IsPrime(17)
	//fmt.Println("IsPrime : ", isPrime)
	//
	//// factorial
	//fmt.Println("Factorial : ", GetFactorial(5))
	//fmt.Println("Factorial using recursion : ", GetFactorialUsingRecursion(5))
	//
	//// word count
	//fmt.Println("Word count : ", GetWordCount("hello get the word count here"))
	//
	//// random
	//fmt.Println(GetRandomNumber(5, 10))
	//fmt.Println(GetRandomNumber(5, 10))

	// matching
	ob := NewOrderBook()

	ob.AddOrder(&Order{ID: 1, Type: SELL, Price: 100.0, Quantity: 10})
	ob.AddOrder(&Order{ID: 2, Type: SELL, Price: 99.0, Quantity: 5})
	ob.AddOrder(&Order{ID: 3, Type: BUY, Price: 101.0, Quantity: 12})
	ob.AddOrder(&Order{ID: 4, Type: BUY, Price: 98.0, Quantity: 7})

	fmt.Println("Remaining orders:")
	fmt.Println("Asks:")
	for _, a := range ob.Asks {
		fmt.Printf("ASK: %d @ %.2f\n", a.Quantity, a.Price)
	}
	fmt.Println("Bids:")
	for _, b := range ob.Bids {
		fmt.Printf("BID: %d @ %.2f\n", b.Quantity, b.Price)
	}

	// my matching
	myOb := NewOrderBookWithRbt()
	myOb.InsertOrder(&Order{ID: 1, Type: SELL, Price: 100.0, Quantity: 10})
	myOb.InsertOrder(&Order{ID: 2, Type: SELL, Price: 99.0, Quantity: 5})
	myOb.InsertOrder(&Order{ID: 3, Type: BUY, Price: 101.0, Quantity: 12})
	myOb.InsertOrder(&Order{ID: 4, Type: BUY, Price: 98.0, Quantity: 7})

	// linked list
	l3 := &LinkedNode{Value: 4, Next: nil}
	l2 := &LinkedNode{Value: 3, Next: l3}
	l1 := &LinkedNode{Value: 2, Next: l2}
	head := &LinkedNode{Value: 1, Next: l1}

	IterateLinkedList(head)
	IsCyclePresent(head)

	newHead := RevertLinkedList(head)
	IterateLinkedList(newHead)
	IsCyclePresent(newHead)

	ll3 := &LinkedNode{Value: 4, Next: nil}
	ll2 := &LinkedNode{Value: 3, Next: ll3}
	ll1 := &LinkedNode{Value: 2, Next: ll2}
	headl := &LinkedNode{Value: 1, Next: ll1}
	ll3.Next = ll1
	IsCyclePresent(headl)

	// merge two sorted arrays
	first := []int{1, 2, 6, 9, 10}
	second := []int{1, 2, 3, 7, 8, 8}
	result := MergeSortedArray(first, second)
	fmt.Println("Merge Sorted Array:", result)

	// rate limiter
	rl := NewUserRateLimiter(10, 10)
	for i := 0; i < 20; i++ {
		fmt.Println("Is request allowed ", i, " ", rl.IsRequestAllowed())
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Retry after 5 seconds")
	for i := 0; i < 5; i++ {
		fmt.Println("Is request allowed ", i, " ", rl.IsRequestAllowed())
	}

}
