package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var doChecksum bool
var doHelp bool
var rootPath string

func doParsing() bool {
	flag.BoolVar(&doChecksum, "c", false, "Calculate file checksums")
	flag.BoolVar(&doHelp, "h", false, "Show usage")
	flag.Parse()
	if doHelp {
		fmt.Printf("Usage of %s: [root]\n", os.Args[0])
		flag.PrintDefaults()
		return false
	}
	if flag.NArg() > 1 {
		fmt.Println("Way too many parameters left unparsed. At maximum you should provide one extra path.")
		return false
	} else if flag.NArg() == 0 {
		var err error
		rootPath, err = os.Getwd()
		if err != nil {
			log.Panic("Could not read current directory. Bye.")
		}
	} else {
		rootPath = os.Args[len(os.Args)-1]
	}
	return true
}

func hashFile(path string) string {
	hash := sha256.New()
	f, err := os.Open(path)
	if err != nil {
		return hashUndefined("?")
	}
	defer f.Close()
	if _, err := io.Copy(hash, f); err != nil {
		return hashUndefined("?")
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func hashUndefined(rep string) string {
	var hash string
	for i := 0; i < 21; i++ {
		hash += rep + "  "
	}
	return hash
}

var lastKnownDir = ""

func visitor(path string, info os.FileInfo, err error) error {
	var hash = ""
	if !info.IsDir() {
		if doChecksum {
			hash = hashFile(path)
		} else {
			hash = hashUndefined(".")
		}
	}
	var dir, file string
	if info.IsDir() {
		dir, file = path+"/", ""
	} else {
		dir, file = filepath.Split(path)
	}
	if dir != lastKnownDir {
		fmt.Println()
		fmt.Println(dir)
		lastKnownDir = dir
	}
	if len(file) > 0 {
		fmt.Printf("%64v %#o %10d %v \n", hash, info.Mode().Perm(), info.Size(), file)
	}
	return nil
}

func main() {
	if doParsing() {
		filepath.Walk(rootPath, visitor)
	}
}
