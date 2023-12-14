package main

import (
	"bytes"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

// setupBench is a non-interactive version of main that can be used for
// benchmarking. Heavily lifted from main(), I just wrote this because it
// seemed easier than running the tests manually one-by-one.
// wrapped in sync.OnceFunc because it changes global state and we only need to
// seed the bloom filter once.
var setupBench = sync.OnceFunc(func() {
	// walk the directory getting files and indexing
	_ = filepath.Walk(".", func(root string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil // we only care about files
		}

		res, err := os.ReadFile(root)
		if err != nil {
			return nil // swallow error
		}

		// don't index binary files by looking for nul byte, similar to how grep does it
		if bytes.IndexByte(res, 0) != -1 {
			return nil
		}

		// only index up to about 5kb
		if len(res) > 5000 {
			res = res[:5000]
		}

		// add the document to the index
		_ = Add(Itemise(Tokenize(string(res))))
		// store the association from what's in the index to the filename, we know its 0 to whatever so this works
		idToFile = append(idToFile, root)
		return nil
	})
})

func BenchmarkMain(b *testing.B) {
	trigramMethod = "jamesrom"
	setupBench()
	for i := 0; i < b.N; i++ {
		Search(Queryise("test"))
	}
}
