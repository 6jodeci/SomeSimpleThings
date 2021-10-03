/*Этот инструмент принимает список файлов из командной строки
//и сжимает по отдельности каждый из файлов, создавая файлы архи-
//вов с теми же именами, но с расширением .gz.
*/
package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var i int = -1
	var file string
	for i, file = range os.Args[1:] {
		wg.Add(1) // for each file, tell the group, that another compression operation is pending
		go func(filename string) {
			compress(filename) // calling a compression operation
			wg.Done()          // and notice of completion
		}(file) // parameter transfer
	}
	wg.Wait() // waiting for all programs performing compression to call wg.Done
	fmt.Printf("Compressed %d files\n", i+1)
}

func compress(filename string) error {
	in, err := os.Open(filename) // open the source file for reading
	if err != nil {
		return fmt.Errorf("filed to open the source file for reading. error: %v", err)
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz") // open a file with a "gz" extension
	if err != nil {
		return fmt.Errorf("failed to open file. error: %v", err)
	}
	defer out.Close()

	gzout := gzip.NewWriter(out) // compress and write the data to an appropriate file
	_, err = io.Copy(gzout, in)  // necessary copying
	gzout.Close()

	return err
}

// go run gzip-compress.go exampledata/example1.txt or go run gzip-compress.go exampledata/*
