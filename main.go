package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kovetskiy/docopt-go"
)

var (
	usage = `CSV gap finder.

Usage:
  csv-gap-finder [options] [<file>] -i <interval> [-d <delimiter>] [-t <timespec>] [-c <column>]
  csv-gap-finder -h | --help

Examples:
  csv-gap-finder timeseries.csv -i 1m -d ";" -t microsecond -c 1
  grep BTCUSDT timeseries.csv | csv-gap-finder -i 1m -d ";l"

Options:
  -i --interval <duration>         Interval between rows. Examples: 1m, 15m, 1d, 30d
  -d --delimiter <delimiter>       CSV delimiter. [default: ,]
  -t --timespec <timespec>         Timespec of timestamps. Examples: second, millisecond, microsecond. [default: second]
  -c --column <index>              Index of timestamp column. [default: 0]
  -h --help                        Show this screen.
  --version                        Show version.`
)

type arguments struct {
	File      string `docopt:"<file>"`
	Interval  string `docopt:"--interval"`
	Delimiter string `docopt:"--delimiter"`
	Timespec  string `docopt:"--timespec"`
	Column    int    `docopt:"--column"`
}

func main() {
	opts, err := docopt.ParseArgs(usage, nil, "")
	if err != nil {
		log.Fatal(err)
	}

	var args arguments
	if err := opts.Bind(&args); err != nil {
		log.Fatal(err)
	}

	var f io.ReadCloser
	if args.File == "" {
		f = os.Stdin
	} else {
		f, err = os.Open(args.File)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer f.Close()

	var durationMultiply time.Duration
	switch args.Timespec {
	case "second":
		durationMultiply = time.Second
	case "millisecond":
		durationMultiply = time.Millisecond
	case "microsecond":
		durationMultiply = time.Microsecond
	}

	delimiter := []byte(args.Delimiter)
	prevTimestamp := 0
	interval, err := time.ParseDuration(args.Interval)
	if err != nil {
		log.Fatal(err)
	}

	writer := os.Stdout

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		columns := bytes.Split(scanner.Bytes(), delimiter)
		if len(columns) < args.Column {
			continue
		}

		ts, err := strconv.Atoi(string(columns[args.Column]))
		if err != nil {
			log.Fatal(err)
		}

		diff := time.Duration(ts-prevTimestamp) * durationMultiply
		if prevTimestamp != 0 && diff > interval {
			fmt.Fprintf(writer, "%d%s%d%s%s\n", ts, args.Delimiter, prevTimestamp, args.Delimiter, diff)
		}

		prevTimestamp = ts
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
