package dictionaries

type FreqDict map[int32]int64

func CreateFreqDictFromLines(lines []string) (dict FreqDict) {
	dict = make(map[int32]int64, 0)
	for _, line := range lines {

		subDict := CreateFreqDictFromLine(line)

		// copying
		for key, value := range subDict {
			dict[key] = value
		}
	}
	return
}

func CreateFreqDictFromLine(line string) (dict FreqDict) {
	dict = make(map[int32]int64, 0)
	for _, symbol := range line {
		dict[symbol]++
	}
	return
}
