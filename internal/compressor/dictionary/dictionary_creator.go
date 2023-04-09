package dictionary

func CreateDictFromLines(lines []string) (dict map[int32]int64) {
	dict = make(map[int32]int64, 0)
	for _, line := range lines {

		subDict := CreateDictFromLine(line)

		// copying
		for key, value := range subDict {
			dict[key] = value
		}
	}
	return
}

func CreateDictFromLine(line string) (dict map[int32]int64) {
	dict = make(map[int32]int64, 0)
	for _, symbol := range line {
		dict[symbol]++
	}
	return
}
