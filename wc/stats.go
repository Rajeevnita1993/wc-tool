package wc

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

type Stats struct {
	LineCount int
	WordCount int
	ByteCount int
}

func CountStats(filename string) (Stats, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Stats{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var stats Stats

	for scanner.Scan() {
		stats.WordCount++
	}

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		stats.LineCount++
	}

	file.Seek(0, 0)
	fileInfo, _ := file.Stat()
	stats.ByteCount = int(fileInfo.Size())

	return stats, nil
}

func CountStatsFromReader(reader io.Reader) (Stats, error) {
	var stats Stats

	// Read the data from the reader and store it in a buffer
	buffer := new(bytes.Buffer)
	teeReader := io.TeeReader(reader, buffer)

	// Count lines
	stats.LineCount = CountLines(teeReader)

	// Count bytes
	stats.ByteCount = buffer.Len()

	// Count words
	stats.WordCount = CountWords(buffer)

	return stats, nil
}

// CountLines counts the number of lines in the given buffer.
func CountLines(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

// CountWords counts the number of words in the given buffer.
func CountWords(buffer *bytes.Buffer) int {
	wordCount := 0
	scanner := bufio.NewScanner(buffer)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordCount++
	}
	return wordCount
}
