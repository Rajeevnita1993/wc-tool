package cmd

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Rajeevnita1993/wc-tool/wc"
)

func Execute() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-c | -w | -l | -m] <filename>\n", os.Args[0])
		flag.PrintDefaults()
	}

	countBytes := flag.Bool("c", false, "Count bytes")
	countWords := flag.Bool("w", false, "Count words")
	countLines := flag.Bool("l", false, "Count lines")
	countChars := flag.Bool("m", false, "Count characters")

	flag.Parse()

	if flag.NArg() > 1 {
		flag.Usage()
		return
	}

	// check if a filename is provided or read from standard input
	var reader io.Reader
	var filename string

	if flag.NArg() == 1 {
		// open file for reading
		filename = flag.Arg(0)
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		defer file.Close()
		reader = file
	} else {
		reader = os.Stdin
	}

	var countType string

	if *countBytes {
		countType = "Bytes"
	} else if *countChars {
		countType = "Chars"
	} else if *countWords {
		countType = "Words"
	} else if *countLines {
		countType = "Lines"
	} else {
		countType = "All"
		*countLines, *countBytes, *countWords = true, true, true
	}

	stats, err := wc.CountStatsFromReader(reader)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	switch countType {
	case "Bytes", "Chars":
		fmt.Printf("%8d %s\n", stats.ByteCount, filename)
	case "Words":
		fmt.Printf("%8d %s\n", stats.WordCount, filename)
	case "Lines":
		fmt.Printf("%8d %s\n", stats.LineCount, filename)
	case "All":
		fmt.Printf("%4d %4d %4d %s\n", stats.LineCount, stats.WordCount, stats.ByteCount, filename)

	}

}
