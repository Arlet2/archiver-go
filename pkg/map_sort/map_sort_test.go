package map_sort

import (
	"reflect"
	"testing"
)

func TestSortingMapByKeys(t *testing.T) {
	t.Log("Start checking sorting of maps by keys")
	{
		testID := 0

		t.Logf("\tTest %d: check sorting map by keys in ascending order (ints)", testID)
		{
			m := map[int]string{5: "lol", 2: "kek", 0: "sd"}
			orderedKeys := []int{0, 2, 5}

			keys := SortMapByKeys(m, true)

			if !reflect.DeepEqual(keys, orderedKeys) {
				t.Logf("\tFail on test %d. Incorrect sort. Expected: %v but got %v", testID, orderedKeys, keys)
				t.FailNow()
			}

		}
		testID++

		t.Logf("\tTest %d: check sorting map by keys in ascending order (strings)", testID)
		{
			m := map[string]string{"B": "lol", "C": "kek", "D": "sd"}
			orderedKeys := []string{"B", "C", "D"}

			keys := SortMapByKeys(m, true)

			if !reflect.DeepEqual(keys, orderedKeys) {
				t.Logf("\tFail on test %d. Incorrect sort. Expected: %v but got %v", testID, orderedKeys, keys)
				t.FailNow()
			}
		}
		testID++

		t.Logf("\tTest %d: check sorting map by keys in descending order (ints)", testID)
		{
			m := map[int]string{5: "lol", 2: "kek", 0: "sd"}
			orderedKeys := []int{5, 2, 0}

			keys := SortMapByKeys(m, false)

			if !reflect.DeepEqual(keys, orderedKeys) {
				t.Logf("\tFail on test %d. Incorrect sort. Expected: %v but got %v", testID, orderedKeys, keys)
				t.FailNow()
			}

		}
		testID++

		t.Logf("\tTest %d: check sorting map by keys in descending order (strings)", testID)
		{
			m := map[string]string{"B": "lol", "C": "kek", "D": "sd"}
			orderedKeys := []string{"D", "C", "B"}

			keys := SortMapByKeys(m, false)

			if !reflect.DeepEqual(keys, orderedKeys) {
				t.Logf("\tFail on test %d. Incorrect sort. Expected: %v but got %v", testID, orderedKeys, keys)
				t.FailNow()
			}
		}
		testID++

		t.Logf("\tTest %d: check sorting empty map", testID)
		{
			m := map[string]string{}

			keys := SortMapByKeys(m, true)

			if !reflect.DeepEqual(keys, []string{}) {
				t.Logf("\tFail on test %d. Incorrect sort. Expected empty slice", testID)
				t.FailNow()
			}
		}
		testID++
	}
}

func TestSortMapByValues(t *testing.T) {
	t.Log("Start checking sorting of maps by values")
	{
		testID := 0

		t.Logf("\tTest %d: check sorting map by values in ascending order (ints)", testID)
		{
			m := map[int]int{5: 5, 2: -5, 0: 3}
			orderedKeys := []int{2, 0, 5}

			keys := SortMapByValues(m, true)

			if !reflect.DeepEqual(keys, orderedKeys) {
				t.Logf("\tFail on test %d. Incorrect sort. Expected: %v but got %v", testID, orderedKeys, keys)
				t.FailNow()
			}

		}
		testID++

		t.Logf("\tTest %d: check sorting map by values in ascending order (strings)", testID)
		{
			m := map[int]string{2: "clol", 3: "akek", 0: "bsd"}
			orderedKeys := []int{3, 0, 2}

			keys := SortMapByValues(m, true)

			if !reflect.DeepEqual(keys, orderedKeys) {
				t.Logf("\tFail on test %d. Incorrect sort. Expected: %v but got %v", testID, orderedKeys, keys)
				t.FailNow()
			}
		}
		testID++

		t.Logf("\tTest %d: check sorting map by values in descending order (ints)", testID)
		{
			m := map[int]int{5: 5, 2: -5, 0: 3}
			orderedKeys := []int{5, 0, 2}

			keys := SortMapByValues(m, false)

			if !reflect.DeepEqual(keys, orderedKeys) {
				t.Logf("\tFail on test %d. Incorrect sort. Expected: %v but got %v", testID, orderedKeys, keys)
				t.FailNow()
			}

		}
		testID++

		t.Logf("\tTest %d: check sorting map by keys in descending order (strings)", testID)
		{
			m := map[string]string{"B": "alol", "C": "ckek", "D": "bsd"}
			orderedKeys := []string{"C", "D", "B"}

			keys := SortMapByValues(m, false)

			if !reflect.DeepEqual(keys, orderedKeys) {
				t.Logf("\tFail on test %d. Incorrect sort. Expected: %v but got %v", testID, orderedKeys, keys)
				t.FailNow()
			}
		}
		testID++

		t.Logf("\tTest %d: check sorting empty map", testID)
		{
			m := map[string]string{}

			keys := SortMapByKeys(m, true)

			if !reflect.DeepEqual(keys, []string{}) {
				t.Logf("\tFail on test %d. Incorrect sort. Expected empty slice", testID)
				t.FailNow()
			}
		}
		testID++
	}
}
