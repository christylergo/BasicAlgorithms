package main

import (
	"fmt"
	"math"
)

type binaryTreeNode struct {
	parent     *binaryTreeNode
	leftChild  *binaryTreeNode
	rightChild *binaryTreeNode
	// name string
	value int
}

type binaryTree struct {
	root *binaryTreeNode
}

func main() {
	var nums = []int{2, 5, 9, 1, 2, 5, 17, 9, 1, 0, 7, 6, 3, 8, 4, 0, 7, 6, 3, 8, 4}
	tree := binaryTree{root: nil}
	tree.initBinaryTree(nums)
	tree.invertBinaryTree(tree.root)
	tree.breadthFirstSearch()

}

func (bt *binaryTree) initBinaryTree(nums []int) {
	if bt.root != nil {
		return
	}
	bt.root = &binaryTreeNode{parent: nil, leftChild: nil, rightChild: nil, value: nums[0]}
	insertNode(bt.root, nums, 0)
}

func insertNode(node *binaryTreeNode, nums []int, i int) {
	if n := len(nums) - 1; 2*i+2 <= n {
		node.leftChild = &binaryTreeNode{parent: node, leftChild: nil, rightChild: nil, value: nums[2*i+1]}
		node.rightChild = &binaryTreeNode{parent: node, leftChild: nil, rightChild: nil, value: nums[2*i+2]}
		insertNode(node.leftChild, nums, 2*i+1)
		insertNode(node.rightChild, nums, 2*i+2)
	} else if 2*i+1 == n {
		node.leftChild = &binaryTreeNode{parent: node, leftChild: nil, rightChild: nil, value: nums[2*i+1]}
	}
}

func (bt *binaryTree) invertBinaryTree(node *binaryTreeNode) {
	if node == nil {
		return
	}
	bt.invertBinaryTree(node.leftChild)
	bt.invertBinaryTree(node.rightChild)
	node.leftChild, node.rightChild = node.rightChild, node.leftChild
}

func (bt binaryTree) breadthFirstSearch() *binaryTreeNode {
	var level, index int
	var nodeQueue []*binaryTreeNode
	nodeQueue = append(nodeQueue, bt.root)
	for ; ; level++ {
		if index >= len(nodeQueue) {
			break
		}
		for ; index < int(math.Pow(2, float64(level+1))-1); index++ {
			if index < len(nodeQueue) {
				tempNode := nodeQueue[index]
				fmt.Print(tempNode.value, "; ")
				if tempNode.leftChild != nil {
					nodeQueue = append(nodeQueue, tempNode.leftChild)
				}
				if tempNode.rightChild != nil {
					nodeQueue = append(nodeQueue, tempNode.rightChild)
				}
			} else {
				break
			}
		}

		fmt.Print("\n************************\n")
	}
	return nil
}
