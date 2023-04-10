package coding_tree

import (
	"archiver/internal/compressor/dictionaries"
	"fmt"
	"testing"
)

func TestGenerateCodingTree(t *testing.T) {
	t.Log("Start checking creating coding tree")
	{
		testID := 0

		t.Logf("\tTest %d: check creating tree with even count of nodes", testID)
		{
			dict := dictionaries.FreqDict{5: 2, 0: 8, 4: 4, 7: 1} // sorted:  7(1) -> 5(2) -> 4(4) -> 0(8)
			tree := CreateCodingTree(dict)

			expectedTree := CodingTree(node{
				freq: 15,
				leftChild: &node{
					freq: 7,
					leftChild: &node{
						freq: 3,
						leftChild: &node{
							freq:  1,
							value: 7,
						},
						rightChild: &node{
							freq:  2,
							value: 5,
						},
					},
					rightChild: &node{
						freq:  4,
						value: 4,
					},
				},
				rightChild: &node{
					freq:  8,
					value: 0,
				},
			})

			if !IsTreeEquals((*node)(&tree), (*node)(&expectedTree)) {
				t.Logf("\tFail on test %d. Incorrect tree. Expected another. Check trees printing", testID)
				fmt.Println("Got:")
				PrintTree(tree)
				fmt.Println("Expected:")
				PrintTree(expectedTree)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check creating tree with odd count of nodes", testID)
		{
			dict := dictionaries.FreqDict{20: 5, 5: 0, 30: 60} // sorted:  5(0) -> 20(5) -> 30(60)
			tree := CreateCodingTree(dict)

			expectedTree := CodingTree(node{
				freq: 65,
				leftChild: &node{
					freq: 5,
					leftChild: &node{
						freq:  0,
						value: 5,
					},
					rightChild: &node{
						freq:  5,
						value: 20,
					},
				},
				rightChild: &node{
					freq:  60,
					value: 30,
				},
			})

			if !IsTreeEquals((*node)(&tree), (*node)(&expectedTree)) {
				t.Logf("\tFail on test %d. Incorrect tree. Expected another. Check trees printing", testID)
				fmt.Println("Got:")
				PrintTree(tree)
				fmt.Println("Expected:")
				PrintTree(expectedTree)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check creating tree with one node", testID)
		{
			dict := dictionaries.FreqDict{20: 5}
			tree := CreateCodingTree(dict)

			expectedTree := CodingTree(node{
				freq:  5,
				value: 20,
			})

			if !IsTreeEquals((*node)(&tree), (*node)(&expectedTree)) {
				t.Logf("\tFail on test %d. Incorrect tree. Expected another. Check trees printing", testID)
				fmt.Println("Got:")
				PrintTree(tree)
				fmt.Println("Expected:")
				PrintTree(expectedTree)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check creating tree with none node", testID)
		{
			dict := dictionaries.FreqDict{}
			tree := CreateCodingTree(dict)

			expectedTree := CodingTree(node{})

			if !IsTreeEquals((*node)(&tree), (*node)(&expectedTree)) {
				t.Logf("\tFail on test %d. Incorrect tree. Expected another. Check trees printing", testID)
				fmt.Println("Got:")
				PrintTree(tree)
				fmt.Println("Expected:")
				PrintTree(expectedTree)
				t.Fail()
			}

		}
		testID++
	}
}
