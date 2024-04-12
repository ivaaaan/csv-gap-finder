# CSV gap finder

Simple utility that allows to find gaps in ordered CSV file. It scans over each row, and verifies whether gap between two rows is bigger than specified interval.

## Usage

```
Usage:
  csv-gap-finder [options] <file> -i <interval> [-d <delimiter>] [-t <timespec>] [-c <column>]
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
```

# LICENSE
MIT
