package encoding

import (
	"archiver/pkg/map_sort"
	"fmt"
	"sort"
)

type CodingTree Node

type Node struct {
	LeftChild  *Node
	RightChild *Node

	Value int32
	freq  int64
}

func CreateCodingTree(dict FreqDict) CodingTree {

	if len(dict) == 0 {
		return CodingTree{}
	}

	sortedKeys := map_sort.SortMapByValues(dict, true)

	nodes := make([]Node, len(sortedKeys))

	// creating base free nodes tree
	for index := range nodes {
		nodes[index] = Node{LeftChild: nil, RightChild: nil, Value: sortedKeys[index], freq: dict[sortedKeys[index]]}
	}

	for len(nodes) > 1 {
		// merge nodes with the lowest freq
		parentNode := Node{LeftChild: &nodes[0], RightChild: &nodes[1], freq: nodes[0].freq + nodes[1].freq}

		// remove child nodes
		nodes = nodes[2:]

		// add parent nodes
		nodes = append(nodes, parentNode)

		// resort all nodes
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].freq < nodes[j].freq
		})
	}

	return CodingTree(nodes[0])
}

func IsTreeEquals(tree *Node, tree1 *Node) bool {
	if tree == nil && tree1 != nil || tree != nil && tree1 == nil {
		return false
	}

	// mean both of them
	if tree == nil {
		return true
	}

	if tree.freq != tree1.freq || tree.Value != tree1.Value {
		return false
	}

	return IsTreeEquals(tree.LeftChild, tree1.LeftChild) && IsTreeEquals(tree.RightChild, tree1.RightChild)
}

func PrintTree(tree CodingTree) {
	dfsPrint((*Node)(&tree), "", "", "")
	fmt.Println()
}

func dfsPrint(node *Node, rpref string, cpref string, lpref string) {
	if node == nil {
		return
	}

	if node.RightChild != nil {
		dfsPrint(node.RightChild, rpref+"  ", rpref+"/--", rpref+"|  ")
	}

	fmt.Printf(cpref+"(f: %d v: %d)\n", node.freq, node.Value)

	if node.LeftChild != nil {
		dfsPrint(node.LeftChild, lpref+"|  ", lpref+"\\--", lpref+"   ")
	}
}
