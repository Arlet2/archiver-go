package coding_tree

import (
	"archiver/internal/compressor/dictionaries"
	"archiver/pkg/map_sort"
	"fmt"
	"sort"
)

type CodingTree node

type node struct {
	leftChild  *node
	rightChild *node

	value int32
	freq  int64
}

func CreateCodingTree(dict dictionaries.FreqDict) CodingTree {

	if len(dict) == 0 {
		return CodingTree{}
	}

	sortedKeys := map_sort.SortMapByValues(dict, true)

	nodes := make([]node, len(sortedKeys))

	// creating base free nodes tree
	for index := range nodes {
		nodes[index] = node{leftChild: nil, rightChild: nil, value: sortedKeys[index], freq: dict[sortedKeys[index]]}
	}

	for len(nodes) > 1 {
		// merge nodes with the lowest freq
		parentNode := node{leftChild: &nodes[0], rightChild: &nodes[1], freq: nodes[0].freq + nodes[1].freq}

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

func IsTreeEquals(tree *node, tree1 *node) bool {
	if tree == nil && tree1 != nil || tree != nil && tree1 == nil {
		return false
	}

	// mean both of them
	if tree == nil {
		return true
	}

	if tree.freq != tree1.freq || tree.value != tree1.value {
		return false
	}

	return IsTreeEquals(tree.leftChild, tree1.leftChild) && IsTreeEquals(tree.rightChild, tree1.rightChild)
}

func PrintTree(tree CodingTree) {
	dfsPrint((*node)(&tree), "", "", "")
	fmt.Println()
}

func dfsPrint(node *node, rpref string, cpref string, lpref string) {
	if node == nil {
		return
	}

	if node.rightChild != nil {
		dfsPrint(node.rightChild, rpref+"  ", rpref+"/--", rpref+"|  ")
	}

	fmt.Printf(cpref+"(f: %d v: %d)\n", node.freq, node.value)

	if node.leftChild != nil {
		dfsPrint(node.leftChild, lpref+"|  ", lpref+"\\--", lpref+"   ")
	}
}
