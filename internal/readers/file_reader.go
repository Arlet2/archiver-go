package readers

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFile(file *os.File) (lines []string, rErr error) {
	reader := bufio.NewReader(file)
	lines = make([]string, 0)

	for {
		bytes, err := reader.ReadBytes('\n')

		//fmt.Println(string(bytes))
		lines = append(lines, string(bytes))

		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF.")
				break
			}
			rErr = err
			return
		}
	}

	return
}
