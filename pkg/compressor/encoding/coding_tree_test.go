package encoding

import (
	"fmt"
	"testing"
)

func TestGenerateCodingTree(t *testing.T) {
	t.Log("Start checking creating coding tree")
	{
		testID := 0

		t.Logf("\tTest %d: check creating tree with even count of nodes", testID)
		{
			dict := FreqDict{5: 2, 0: 8, 4: 4, 7: 1} // sorted:  7(1) -> 5(2) -> 4(4) -> 0(8)
			tree := CreateCodingTree(dict)

			expectedTree := CodingTree(Node{
				freq: 15,
				LeftChild: &Node{
					freq: 7,
					LeftChild: &Node{
						freq: 3,
						LeftChild: &Node{
							freq:  1,
							Value: 7,
						},
						RightChild: &Node{
							freq:  2,
							Value: 5,
						},
					},
					RightChild: &Node{
						freq:  4,
						Value: 4,
					},
				},
				RightChild: &Node{
					freq:  8,
					Value: 0,
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
			dict := FreqDict{20: 5, 5: 0, 30: 60} // sorted:  5(0) -> 20(5) -> 30(60)
			tree := CreateCodingTree(dict)

			expectedTree := CodingTree(Node{
				freq: 65,
				LeftChild: &Node{
					freq: 5,
					LeftChild: &Node{
						freq:  0,
						Value: 5,
					},
					RightChild: &Node{
						freq:  5,
						Value: 20,
					},
				},
				RightChild: &Node{
					freq:  60,
					Value: 30,
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
			dict := FreqDict{20: 5}
			tree := CreateCodingTree(dict)

			expectedTree := CodingTree(Node{
				freq:  5,
				Value: 20,
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
			dict := FreqDict{}
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
