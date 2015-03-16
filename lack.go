package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"strings"

	"github.com/apg/lack/logfmt"
	"github.com/apg/lack/query"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var format = flag.String("format", "", "format string for output")

func main() {
	var q query.Query

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if flag.NArg() >= 1 {
		q, _ = query.Parse([]byte(flag.Arg(0)))
	}

	reader := bufio.NewReader(os.Stdin)

	line := logfmt.NewLine()

	if *format != "" {
		*format = strings.Replace(*format, "\\t", "\t", -1)
	}

	for {
		l, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		line.Reset(l)
		err = logfmt.Scan(line)
		if q != nil && q.Match(line) {
			if *format != "" {
				if out, err := line.Format([]byte(*format)); err != nil {
					fmt.Fprintf(os.Stderr, "Error: %s", err)
					os.Exit(1)
				} else {
					fmt.Println(out)
				}
			} else {
				fmt.Println(line.Bytes())
			}
		}
	}
}
