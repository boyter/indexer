package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// keeps track of files stored in the index so we can open them to find matches
var idToFile []string

var trigramMethod = ""

func main() {

	flag.StringVar(&trigramMethod, "trigram", "default", "which trigram method should we use [default,merovius,dancantos,jamesrom,ffmiruz]")
	flag.Parse()

	startTime := time.Now()
	// walk the directory getting files and indexing
	//_ = filepath.Walk(".", func(root string, info os.FileInfo, err error) error {
	//	if err != nil {
	//		return err
	//	}
	//
	//	if info.IsDir() {
	//		return nil // we only care about files
	//	}
	//
	//	res, err := os.ReadFile(root)
	//	if err != nil {
	//		return nil // swallow error
	//	}
	//
	//	// don't index binary files by looking for nul byte, similar to how grep does it
	//	if bytes.IndexByte(res, 0) != -1 {
	//		return nil
	//	}
	//
	//	// only index up to about 5kb
	//	if len(res) > 5000 {
	//		res = res[:5000]
	//	}
	//
	//	// add the document to the index
	//	_ = Add(Itemise(Tokenize(string(res))))
	//	// store the association from what's in the index to the filename, we know its 0 to whatever so this works
	//	idToFile = append(idToFile, root)
	//	return nil
	//})

	rand.Seed(1)
	for j := 0; j < 10000; j++ {
		for i := 0; i < BloomSize; i++ {
			bloomFilter = append(bloomFilter, rand.Uint64())
			idToFile = append(idToFile, strconv.Itoa(i))
		}
	}

	endTime := time.Since(startTime)
	fmt.Printf("currentBlockDocumentCount:%v currentDocumentCount:%v currentBlockStartDocumentCount:%v indexTimeSeconds:%v trigramMethod:%v\n", currentBlockDocumentCount, currentDocumentCount, currentBlockStartDocumentCount, endTime.Seconds(), trigramMethod)
	if trigramMethod == "" {
		return
	}

	for _, searchTerm := range []string{"test", "import", "struct", "linux", "logitech", "boyterwashere", "linus", "fuck", "shit", "thisshouldmatchnothing"} {
		startTime = time.Now()
		res := Search(Queryise(searchTerm))
		fmt.Println(fmt.Sprintf("len(%v) ms(%v)", len(res), time.Since(startTime).Milliseconds()))
	}
}

// Given a file and a query try to open the file, then look through its lines
// and see if any of them match something from the query up to a limit
// Note this will return partial matches as if any term matches its considered a match
// and there is no accounting for better matches...
// In other words it's a very dumb way of doing this and probably has horrible runtime
// performance to match
func findMatchingLines(filename string, query string, limit int) []string {
	res, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}

	terms := strings.Fields(strings.ToLower(query))
	var cleanTerms []string
	for _, t := range terms {
		if len(t) >= 3 {
			cleanTerms = append(cleanTerms, t)
		}
	}

	var matches []string
	for i, l := range strings.Split(string(res), "\n") {

		low := strings.ToLower(l)
		found := false
		for _, t := range terms {
			if strings.Contains(low, t) {
				if !found {
					matches = append(matches, fmt.Sprintf("%v. %v", i+1, l))
				}
				found = true
			}
		}

		if len(matches) >= limit {
			return matches
		}
	}

	return matches
}
