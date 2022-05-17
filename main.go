package main

import "fmt"

type bstnode struct {
	value int
	left  *bstnode
	right *bstnode
}

type bst struct {
	root *bstnode
}

func (b *bst) reset() {
	b.root = nil
}

func (b *bst) insert(value int) {
	b.insertRecord(b.root, value)
}

func (b *bst) insertRecord(node *bstnode, value int) *bstnode {
	if b.root == nil {
		b.root = &bstnode{
			value: value,
		}
		return b.root
	}

	if node == nil {
		return &bstnode{value: value}
	}
	if value <= node.value {
		node.left = b.insertRecord(node.left, value)
	}
	if value > node.value {
		node.right = b.insertRecord(node.right, value)
	}
	return node

}

func (b *bst) inorder() {
	fmt.Printf("root is : %d\n", b.root.value)
	b.inorderRecord(b.root)
}

func (b *bst) inorderRecord(node *bstnode) {

	if node != nil {
		if node.right != nil && node.left != nil {
			fmt.Printf("node: %d, left : %d , right : %d\n", node.value, node.left.value, node.right.value)
		} else if node.left != nil {
			fmt.Printf("node: %d, left : %d , right : nil\n", node.value, node.left.value)
		} else if node.right != nil {
			fmt.Printf("node: %d, left : nil , right : %d\n", node.value, node.right.value)
		}

		b.inorderRecord(node.left)
		b.inorderRecord(node.right)
	}
}

func main() {
	bst := &bst{}
	bst.insert(12)
	bst.insert(15)
	bst.insert(14)
	bst.insert(12)
	bst.insert(18)
	bst.insert(11)
	bst.insert(16)
	bst.insert(20)

	bst.inorder()

}
