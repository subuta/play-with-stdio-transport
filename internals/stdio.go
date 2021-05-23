package internals

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/subuta/play-with-stdio-transport/pkg/parser"
)

// LoopRead
// ReadString solution for reading text from stdio
// SEE: https://stackoverflow.com/a/47481035/9998350
func LoopRead(parser parser.EventParser) error {
	r := bufio.NewReader(os.Stdin)
	for {
		line, err := r.ReadString('\n')
		if len(line) == 0 && err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		line = strings.TrimSuffix(line, "\n")

		parser.OnData(line)

		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	return nil
}