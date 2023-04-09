package readers

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFile(file *os.File) (lines []string, rErr error) {
	reader := bufio.NewReader(file)
	lines = make([]string, 0)

	for {
		bytes, err := reader.ReadBytes('\n')

		//fmt.Println(string(bytes))
		lines = append(lines, string(bytes)+"\n")

		if err != nil {
			if err == io.EOF {
				fmt.Println("EOF.")
				break
			}
			rErr = err
			return
		}
	}

	// Remove \n in last line because it's got EOF instead
	lines[len(lines)-1] = strings.Replace(lines[len(lines)-1], "\n", "", -1)

	return
}
