package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/apg/lack/logfmt"
	"github.com/apg/lack/query"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

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

	for {
		l, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}

		line.Reset(l)
		err = logfmt.Scan(line)
		if q != nil && q.Match(line) {
			fmt.Fprintf(os.Stdout, "%s", line.Bytes())
		}
	}
}
