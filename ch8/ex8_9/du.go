package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

type rootInfo struct {
	root string
	size int64
}

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Traverse the file tree.
	fileSizes := make(chan *rootInfo)
	nfiles := make(map[string]int64)
	nbytes := make(map[string]int64)
	var n sync.WaitGroup

	for _, root := range roots {
		nfiles[root] = 0
		nbytes[root] = 0
		n.Add(1)
		go walkDir(root, root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()
	// Print the results periodically
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
loop:
	for {
		select {
		case info, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles[info.root]++
			nbytes[info.root] += info.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}
	printDiskUsage(roots, nfiles, nbytes) // final totals
}

func printDiskUsage(roots []string, nfiles, nbytes map[string]int64) {
	for _, root := range roots {
		fmt.Printf("%s:\t%d files %.1f GB\n", root, nfiles[root], float64(nbytes[root])/1e9)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(root string, dir string, n *sync.WaitGroup, fileSizes chan<- *rootInfo) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Add(1)
			go walkDir(root, subdir, n, fileSizes)
		} else if info, err := entry.Info(); err == nil {
			fileSizes <- &rootInfo{root, info.Size()}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []fs.DirEntry {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
