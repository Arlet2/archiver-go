package writers

import (
	"bufio"
	"os"
)

func WriteFile(file *os.File, lines []string) error {
	writer := bufio.NewWriter(file)

	err := writeHeader(writer)

	if err != nil {
		return err
	}

	for _, line := range lines {
		_, err = writer.Write([]byte(line))

		if err != nil {
			return err
		}
	}
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func writeHeader(writer *bufio.Writer) error {
	_, err := writer.Write([]byte("EVA"))

	return err
}
