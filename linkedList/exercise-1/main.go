package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

func (list *linkedList) append(data int) {
	// Create a new node that will be added to the linklist
	newNode := &node{
		data: data,
	}
	// If the linked list is empty, set both head and tail to the new node
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		// If the linked list is not empty, add the new node to the end
		// Update the next reference of the current tail to point to the new node
		list.tail.next = newNode
		// Update the tail to be the new node
		list.tail = newNode
	}
	// Increase the length of the link list by one
	list.length++

}

func (list *linkedList) remove(data int) {
	// if it is head remove
	if list.head.data == data {
		list.head = list.head.next
		list.length--
		return
	}
	// Start from the head of the linked list
	currentNode := list.head

	// Iterate through the linked list until the end is reached
	for currentNode != nil {
		// If the next node is equal to the data and it is the tail, remove
		if currentNode.next == list.tail && list.tail.data == data {
			list.tail = currentNode // Update the tail to be the current node, effectively removing the current tail
			list.length--
			break // Exit the loop after removing the node
		}

		// If the next node is equal to the data, remove
		if currentNode.next.data == data {
			// Update the next reference of the current node to skip the next node
			currentNode.next = currentNode.next.next
			list.length--
			break
		}
		// Move to the next node in the link list
		currentNode = currentNode.next
	}
}

// contains checks if the linked list contains a node with the specified data.
// It returns true if found, otherwise false.
func (list *linkedList) contains(data int) bool {
	// Check if the data is in the head or tail of the linked list
	if list.head.data == data || list.tail.data == data {
		// Data found in either the head or tail, return true
		return true
	}

	// Start from the head of the linked list
	currentNode := list.head

	// Iterate through the linked list until the end is reached
	for currentNode.next != nil {
		// If the current node's data is equal to the specified data, return true
		if currentNode.data == data {
			return true
		}

		// Move to the next node in the linked list
		currentNode = currentNode.next
	}

	// Data not found in the linked list, return false
	return false
}

func (list *linkedList) get(data int) *node {
	if list.head.data == data {
		return list.head
	}
	if list.tail.data == data {
		return list.tail
	}

	currentNode := list.head
	for currentNode.next != nil {
		if currentNode.data == data {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return nil
}

func main() {
	ll := &linkedList{}
	ll.append(1)
	ll.append(2)
	ll.append(3)
	ll.append(4)
	ll.append(5)
	ll.append(6)
	ll.append(7)
	ll.append(8)

	ll.remove(4)
	ll.remove(8)

	fmt.Printf("Contains 4 = %t \n", ll.contains(4))
	fmt.Printf("Contains 7 = %t \n", ll.contains(7))

	n := ll.get(3)
	if n != nil {
		fmt.Printf("Found Node: %d\n", n.data)
	}

	n1 := ll.get(10)
	if n != nil {
		fmt.Printf("No node containing %d \n", n1.data)
	}

}
