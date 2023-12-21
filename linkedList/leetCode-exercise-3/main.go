// Find middle Node

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode // Reference to the next node in the linked list
}

func middleNode(head *ListNode) *ListNode {
	// Count the number of the nodes
	var numberOfNodes int = 1
	for current := head; current.Next != nil; numberOfNodes++ {
		current = current.Next
	}

	// Calculate the index of the middle node
	middleIndex := numberOfNodes / 2

	// Traverse to the middle node in the list
	var middleNode *ListNode = head
	for i := 0; i < middleIndex; i++ {
		middleNode = middleNode.Next
	}
	return middleNode
}

func main() {
	// Create a linked list: 1 -> 2 -> 3 -> 4 -> 5
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	middle := middleNode(head)
	fmt.Println("Middle Node Value: ", middle.Val)
}
