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

			expectedTree := CodingTree(Node{
				freq: 15,
				leftChild: &Node{
					freq: 7,
					leftChild: &Node{
						freq: 3,
						leftChild: &Node{
							freq:  1,
							value: 7,
						},
						rightChild: &Node{
							freq:  2,
							value: 5,
						},
					},
					rightChild: &Node{
						freq:  4,
						value: 4,
					},
				},
				rightChild: &Node{
					freq:  8,
					value: 0,
				},
			})

			if !IsTreeEquals((*Node)(&tree), (*Node)(&expectedTree)) {
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

			expectedTree := CodingTree(Node{
				freq: 65,
				leftChild: &Node{
					freq: 5,
					leftChild: &Node{
						freq:  0,
						value: 5,
					},
					rightChild: &Node{
						freq:  5,
						value: 20,
					},
				},
				rightChild: &Node{
					freq:  60,
					value: 30,
				},
			})

			if !IsTreeEquals((*Node)(&tree), (*Node)(&expectedTree)) {
				t.Logf("\tFail on test %d. Incorrect tree. Expected another. Check trees printing", testID)
				fmt.Println("Got:")
				PrintTree(tree)
				fmt.Println("Expected:")
				PrintTree(expectedTree)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check creating tree with one Node", testID)
		{
			dict := dictionaries.FreqDict{20: 5}
			tree := CreateCodingTree(dict)

			expectedTree := CodingTree(Node{
				freq:  5,
				value: 20,
			})

			if !IsTreeEquals((*Node)(&tree), (*Node)(&expectedTree)) {
				t.Logf("\tFail on test %d. Incorrect tree. Expected another. Check trees printing", testID)
				fmt.Println("Got:")
				PrintTree(tree)
				fmt.Println("Expected:")
				PrintTree(expectedTree)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check creating tree with none Node", testID)
		{
			dict := dictionaries.FreqDict{}
			tree := CreateCodingTree(dict)

			expectedTree := CodingTree(Node{})

			if !IsTreeEquals((*Node)(&tree), (*Node)(&expectedTree)) {
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
