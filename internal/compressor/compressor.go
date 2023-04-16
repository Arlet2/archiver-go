package compressor

import (
	"archiver/internal/compressor/encoding"
	"strconv"
	"strings"
)

const oneSymbolSize = 10 // до 0xFFFF (16 двоичных цифр)

func Compress(lines []string, dict encoding.CodingDict) (nlines []string, err error) {
	var buf string
	var size int

	nlines = make([]string, 0)

	for _, line := range lines {
		//todo: оптимизировать, сделать создание буфера на ходу и добавить слияние двух строк (хвост и новую строку)
		//todo: доделать replaceBySymbols
		if float64(len(line)) > 1.5*float64(len(dict)) || true {
			line = replaceByDict(line, dict)
		} else {
			line = replaceBySymbols(line, dict)
		}

		// symbol conversion
		buf = ""

		for len(line) != 0 {
			if len(line) > oneSymbolSize {
				size = oneSymbolSize
			} else {
				size = len(line)
			}

			newSymbol, err := strconv.ParseUint(line[:size], 2, 32)

			if err != nil {
				panic("In compressing: " + err.Error())
			}

			buf += string(uint32(newSymbol))
			line = line[size:]
		}

		nlines = append(nlines, buf)
	}

	return
}

func replaceByDict(line string, dict encoding.CodingDict) string {
	for key, value := range dict {
		line = strings.ReplaceAll(line, string(key), value)
	}
	return line
}

// todo: дофиксить
func replaceBySymbols(line string, dict encoding.CodingDict) string {
	for i := 0; i < len(line); i++ {
		newSymbol := dict[int32(line[i])]
		line = strings.ReplaceAll(line, string(line[i]), newSymbol)

		i += len(newSymbol) - 1
	}

	return line
}
