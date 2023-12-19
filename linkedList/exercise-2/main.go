package main

import (
	"fmt"
)

type node struct {
	data int
	prev *node
	next *node
}

type doublyLinkedList struct {
	head   *node
	tail   *node
	length int
}

func (list *doublyLinkedList) append(data int) {
	newNode := &node{
		data: data,
		prev: nil,
		next: nil,
	}
	// If the doubly linked list is empty, set both head and tail to the new node
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		// If the doubly linked list is not empty, add the new node to the end
		// Update the next reference of the current tail to point to the new node
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode // Update the tail to be the new node
	}

	list.length++ // Increase the length of the doubly linked list by one
}

func (list *doublyLinkedList) remove(data int) {
	currentNode := list.head

	// Iterate through the doubly linked list until the end is reached
	for currentNode != nil {
		// If the current node's data is equal to the specified data, remove
		if currentNode.data == data {
			// If the node to be removed is the head
			if currentNode.prev == nil {
				list.head = currentNode.next
				if list.head != nil {
					list.head.prev = nil
				} else {
					// If the head is nil, set the tail to nil as well
					list.tail = nil
				}
			} else {
				// If the node to be removed is not the head
				currentNode.prev.next = currentNode.next
				// If the node to be removed is not the tail
				if currentNode.next != nil {
					currentNode.next.prev = currentNode.prev
				} else {
					// If the node to be removed is the tail
					list.tail = currentNode.prev
				}
			}
			list.length--
			break
		}
		currentNode = currentNode.next // Move to the next node in the doubly linked list
	}
}

func (list *doublyLinkedList) contains(data int) bool {
	currentNode := list.head

	for currentNode != nil {
		if currentNode.data == data {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func main() {
	mylist := &doublyLinkedList{}

	mylist.append(10)
	mylist.append(20)
	mylist.append(30)
	mylist.append(40)
	mylist.append(50)
	mylist.append(60)
	mylist.append(70)
	mylist.append(80)

	fmt.Println("Initial doubly link list")
	displaylist(mylist)

	mylist.remove(10)
	mylist.remove(50)
	mylist.remove(70)

	fmt.Println("\nDoubly Linked List after Removing 10,50,70:")
	displaylist(mylist)

	result := mylist.contains(30)
	fmt.Println("\nDoes Doubly Linked List contain 30?", result)

	result = mylist.contains(70)
	fmt.Println("Does Doubly Linked List contain 70?", result)

}

func displaylist(list *doublyLinkedList) {
	currentNode := list.head
	for currentNode != nil {
		fmt.Printf("%d <->", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("nil")
}
