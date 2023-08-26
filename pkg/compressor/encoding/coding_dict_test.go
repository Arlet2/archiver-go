package encoding

import (
	"reflect"
	"testing"
)

func TestCreateCodingDict(t *testing.T) {
	t.Log("Start checking creating coding dict")
	{
		testID := 0

		t.Logf("\tTest %d: check simple creating coding dict", testID)
		{
			freqDict := FreqDict(map[int32]int64{'a': 5, 'f': 3, 'y': 50, 'p': 4})

			tree := CreateCodingTree(freqDict)

			dict := CreateCodingDict(tree)

			expectedDict := CodingDict(map[int32]string{'a': "00", 'f': "010", 'y': "1", 'p': "011"})

			if !reflect.DeepEqual(dict, expectedDict) {
				t.Logf("\tFail on test %d. Expected another dict. Expected: %v. Got: %v",
					testID, expectedDict, dict)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check simple creating coding dict with empty tree", testID)
		{
			dict := CreateCodingDict(CodingTree{})

			expectedDict := CodingDict(map[int32]string{})

			if !reflect.DeepEqual(dict, expectedDict) {
				t.Logf("\tFail on test %d. Expected another dict. Expected: %v. Got: %v",
					testID, expectedDict, dict)
				t.Fail()
			}

		}
		testID++
	}

}
