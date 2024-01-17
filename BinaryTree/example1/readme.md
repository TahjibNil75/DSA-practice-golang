link : https://www.youtube.com/watch?v=D3XlKCii7L4&list=PLve39GJ2D71xX0Ham0WoPaYfl8oTzZfN6&index=19


    tarverse rule: left -> itslef -> right

    int < node : Traverse the left subtree.
    int > node : Traverse the left subtree.

```go
func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.left)  // Traverse the left subtree.
	sb.WriteString(fmt.Sprintf("%s ", root)) // Visit the current node.
	b.inOrderTraversalByNode(sb, root.right) // Traverse the right subtree.
}
```


BST Remove()

1. Method 1:
            Replace node with smallest value on right subtree, then try to delete that value on right subtree. If there is not a right subtree, just return left subtree
2. Method 2:
            Replace node with largest value on left subtree, then try to delete that value on left subtree. If there is not a left subtree, just return right subtree




