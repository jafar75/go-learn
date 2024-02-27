// Using map to struct pointers in order to fulfill the requirement of Exercise 1.4

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Duplicates struct {
	Count int
	FileNames []string
}

func main() {
	counts := make(map[string]*Duplicates)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.Count > 1 {
			fmt.Printf("%d\t%s\t%v\n", n.Count, line, n.FileNames)
		}
	}
}
func countLines(f *os.File, counts map[string]*Duplicates) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		entry, ok := counts[input.Text()];
		if ok {
			entry.Count ++;
			if !slices.Contains(entry.FileNames, f.Name()) {
				entry.FileNames = append(entry.FileNames, f.Name());
			}
		} else {
			entry := Duplicates{
				Count: 1,
				FileNames: []string{f.Name()},
			}
			counts[input.Text()] = &entry;
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
