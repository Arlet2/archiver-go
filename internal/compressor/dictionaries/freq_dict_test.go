package dictionaries

import (
	"reflect"
	"testing"
)

func TestCreateFreqDictFromLine(t *testing.T) {
	t.Log("Start checking generating freq dict from line")
	{
		testID := 0

		t.Logf("\tTest %d: check simple generating freq dict", testID)
		{
			line := "aabb{}[[["

			dict := CreateFreqDictFromLine(line)

			expectedDict := FreqDict(map[int32]int64{'a': 2, 'b': 2, '{': 1, '}': 1, '[': 3})

			if !reflect.DeepEqual(dict, expectedDict) {
				t.Logf("\tFail on test %d. Expected another dict. Expected: %v. Got: %v",
					testID, expectedDict, dict)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check simple generating freq dict with empty line", testID)
		{
			line := ""

			dict := CreateFreqDictFromLine(line)

			expectedDict := FreqDict(map[int32]int64{})

			if !reflect.DeepEqual(dict, expectedDict) {
				t.Logf("\tFail on test %d. Expected another dict. Expected: %v. Got: %v",
					testID, expectedDict, dict)
				t.Fail()
			}

		}
		testID++
	}
}

func TestCreateFreqDictFromLines(t *testing.T) {
	t.Log("Start checking generating freq dict from lines")
	{
		testID := 0

		t.Logf("\tTest %d: check simple generating freq dict", testID)
		{
			lines := []string{"aabb{}[[[", "abbf"}

			dict := CreateFreqDictFromLines(lines)

			expectedDict := FreqDict(map[int32]int64{'a': 3, 'b': 4, '{': 1, '}': 1, '[': 3, 'f': 1})

			if !reflect.DeepEqual(dict, expectedDict) {
				t.Logf("\tFail on test %d. Expected another dict. Expected: %v. Got: %v",
					testID, expectedDict, dict)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check simple generating freq dict with empty lines", testID)
		{
			var lines []string

			dict := CreateFreqDictFromLines(lines)

			expectedDict := FreqDict(map[int32]int64{})

			if !reflect.DeepEqual(dict, expectedDict) {
				t.Logf("\tFail on test %d. Expected another dict. Expected: %v. Got: %v",
					testID, expectedDict, dict)
				t.Fail()
			}

		}
		testID++
	}
}
