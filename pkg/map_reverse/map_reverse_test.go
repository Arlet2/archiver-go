package map_reverse

import (
	"reflect"
	"testing"
)

func TestReverseMap(t *testing.T) {
	t.Log("Start checking reversion of map")
	{
		testID := 0

		t.Logf("\tTest %d: check reversion of simple map", testID)
		{
			m := ReverseMap(map[string]int{"A": 2, "B": 0, "C": 3})

			expectedMap := map[int]string{2: "A", 0: "B", 3: "C"}

			if !reflect.DeepEqual(m, expectedMap) {
				t.Logf("\tFail on test %d. Expected another dict. Expected: %v. Got: %v",
					testID, m, expectedMap)
				t.Fail()
			}

		}
		testID++

		t.Logf("\tTest %d: check reversion of empty map", testID)
		{
			m := ReverseMap(map[string]int{})

			expectedMap := map[int]string{}

			if !reflect.DeepEqual(m, expectedMap) {
				t.Logf("\tFail on test %d. Expected another dict. Expected: %v. Got: %v",
					testID, m, expectedMap)
				t.Fail()
			}

		}
		testID++
	}
}
