package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

const helpStr = `version: 0.0.1, statistics cvs column range

example: %s -in=input.csv -out=output.csv -column=3 -value=8888.88 -delta=6.66 -border=true
         will statistics 3'd column value, range [8888.88-6.66, 8888.88+6.66]

`

var op = &OP{
	in:     "input.csv",
	out:    "output.csv",
	column: 0,
	value:  0,
	delta:  0,
	stdout: true,
	border: true,
}

type OP struct {
	in     string
	out    string
	column int
	value  float64
	delta  float64
	stdout bool
	border bool
	help   bool
}

func (c *OP) Valid() bool {
	if c.column <= 0 || len(c.in) == 0 || op.help {
		return false
	} else {
		return true
	}
}

func (c *OP) InValid() bool {
	return !c.Valid()
}

func (c *OP) GetInputs() ([][]string, error) {
	f, err := os.OpenFile(op.in, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return csv.NewReader(f).ReadAll()
}

func (c *OP) Statistics(records [][]string) ([][]string, error) {
	var result [][]string
	for idx, line := range records {
		if len(line) < op.column {
			return nil, fmt.Errorf("column is to large")
		}
		item := line[op.column-1]
		if len(item) <= 0 {
			continue
		}
		value, err := strconv.ParseFloat(line[op.column-1], 64)
		if err != nil {
			return nil, fmt.Errorf("line %d column %d convert to float failed, %v", idx+1, op.column, err)
		}
		min := op.value - op.delta
		max := op.value + op.delta
		if op.border && min <= value && value <= max {
			result = append(result, line)
		} else if min < value && value < max {
			result = append(result, line)
		}
	}
	return result, nil
}

func (c *OP) OutputFile(records [][]string) error {
	f, err := os.OpenFile(op.out, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	return csv.NewWriter(f).WriteAll(records)
}

func (c *OP) OutputStdout(records [][]string) {
	for idx, line := range records {
		fmt.Printf("%d: %v\n", idx, strings.Join(line, ","))
	}
}

func init() {
	flag.StringVar(&op.in, "in", op.in, fmt.Sprintf("set the input file name, default: %s", op.in))
	flag.StringVar(&op.out, "out", op.out, fmt.Sprintf("set the output file name, default: %s", op.out))
	flag.IntVar(&op.column, "column", op.column, fmt.Sprintf("set column number which will be statistics"))
	flag.Float64Var(&op.value, "value", op.value, fmt.Sprintf("the targe value for statistics"))
	flag.Float64Var(&op.delta, "delta", op.value, fmt.Sprintf("the delta value for statistics"))
	flag.BoolVar(&op.stdout, "stdout", op.stdout, fmt.Sprintf("set results output to stdout, default: %v", op.stdout))
	flag.BoolVar(&op.border, "border", op.border, fmt.Sprintf("value range include border, default: %v", op.border))
	flag.BoolVar(&op.help, "help", op.help, "get help information")
	flag.Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, fmt.Sprintf(helpStr, path.Base(os.Args[0])))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if op.InValid() {
		flag.Usage()
		return
	}
	records, err := op.GetInputs()
	if err != nil {
		log.Fatal(err)
	}
	records, err = op.Statistics(records)
	if err != nil {
		log.Fatal(err)
	}
	if op.stdout {
		op.OutputStdout(records)
	}
	if len(op.out) > 0 {
		err = op.OutputFile(records)
		if err != nil {
			log.Fatal(err)
		}
	}
}
