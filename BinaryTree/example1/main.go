// Package main is the entry point for the program.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// node represents a node in a binary search tree.
type node struct {
	value int   // The value of the node.
	left  *node // The left child of the node.
	right *node // The right child of the node.
}

// bst represents a binary search tree.
type bst struct {
	root *node // The root of the binary search tree.
	len  int   // The number of nodes in the binary search tree.
}

// String method returns the string representation of a node.
func (n node) String() string {
	return strconv.Itoa(n.value)
}

// inOrderTraversalByNode performs in-order traversal starting from a given node.
func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.left)  // Traverse the left subtree.
	sb.WriteString(fmt.Sprintf("%s ", root)) // Visit the current node.
	b.inOrderTraversalByNode(sb, root.right) // Traverse the right subtree.
}

// inOrderTraversal performs in-order traversal of the entire binary search tree.
func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root) // Start in-order traversal from the root.
}

// String method returns the string representation of the binary search tree.
func (b bst) String() string {
	var sb strings.Builder
	b.inOrderTraversal(&sb)
	return strings.TrimSpace(sb.String())
}

// ADD Method --------------------------------------------------------------------------------------------------------------------------------------------------------------------

// addByNode recursively adds a value to the binary search tree starting from a given node.
func (b *bst) addByNode(root *node, value int) *node {
	if root == nil {
		return &node{value: value} // Create a new node with the given value if the current node is nil.
	}
	if value < root.value {
		root.left = b.addByNode(root.left, value) // If the value is less than the current node, traverse the left subtree.
	} else {
		root.right = b.addByNode(root.right, value) // If the value is greater than or equal to the current node, traverse the right subtree.
	}
	return root // Return the modified node.
}

// add adds a value to the binary search tree.
func (b *bst) add(value int) {
	b.root = b.addByNode(b.root, value) // Call addByNode to recursively add the value to the BST.
	b.len++                             // Increment the length of the BST.
}

// --------------------------------------------------------------------------------------------------------------------------------------------------------------------

// Search By Node --------------------------------------------------------------------------------------------------------------------------------------------------------------------

func (b bst) searchByNode(root *node, value int) (*node, bool) {

	if root == nil {
		return nil, false
	}
	if value == root.value {
		return root, true
	} else if value < root.value {
		return b.searchByNode(root.left, value)
	} else if value > root.value {
		return b.searchByNode(root.right, value)
	}
	return nil, false
}

func (b bst) search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}

// --------------------------------------------------------------------------------------------------------------------------------------------------------------------

// Remove Method --------------------------------------------------------------------------------------------------------------------------------------------------------------------

func (b *bst) removeByNode(root *node, value int) *node {

	if root == nil {
		return root
	}

	if value > root.value {
		root.right = b.removeByNode(root.right, value)
	} else if value < root.value {
		root.left = b.removeByNode(root.left, value)
	} else { // method 2 appling
		if root.left == nil {
			return root.right
		} else {
			temp := root.left // temp value consist of max value of the left subtree
			for temp.right != nil {
				temp = temp.right
			}
			root.value = temp.value

			root.left = b.removeByNode(root.left, value)
		}
	}
	return root

}

func (b *bst) remove(value int) {
	b.removeByNode(b.root, value)
}

func main() {
	n := &node{1, nil, nil}
	n.left = &node{0, nil, nil}
	n.right = &node{2, nil, nil}
	b := bst{
		root: n,
		len:  3,
	}
	fmt.Println(b)

	// Test Add method
	fmt.Println("Testing Add method")
	b.add(5)
	b.add(4)
	b.add(6)
	fmt.Println(b)

	fmt.Println("Testing Search method")
	fmt.Println(b.search(6))
	fmt.Println(b.search(3))

	fmt.Println("Testing remove method")
	b.remove(6)
	fmt.Println(b)
	b.remove(3)
	fmt.Println(b)

}
