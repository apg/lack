package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/apg/lack/logfmt"
	"github.com/apg/lack/query"
)

func main() {
	var q query.Matcher

	if len(os.Args) > 1 {
		q, _ = query.Parse([]byte(os.Args[1]))
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		l, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		line, err := logfmt.Scan(l)
		if q != nil && q(line) {
			fmt.Fprintf(os.Stdout, "%s", line.Bytes())
		}
	}
}
