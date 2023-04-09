package writers

import (
	"bufio"
	"os"
)

func WriteFile(file *os.File, lines []string) error {
	writer := bufio.NewWriter(file)

	for _, line := range lines {
		_, err := writer.Write([]byte(line))

		if err != nil {
			return err
		}
	}

	return nil
}
