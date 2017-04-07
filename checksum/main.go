package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

var (
	inputFile = flag.String("f", "", "File for checksum calculating")
	algo      = flag.String("algo", "md5", "Algorithm checksum calculation. Allowed values: md5, sha1, sha256, sha512.")
)

func main() {
	flag.Parse()
	if *inputFile == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}
	f, err := os.Open(*inputFile)
	if err != nil {
		fmt.Printf("Error open file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	var hash hash.Hash
	switch *algo {
	case "md5":
		hash = md5.New()
	case "sha256":
		hash = sha256.New()
	case "sha1":
		hash = sha1.New()
	case "sha512":
		hash = sha512.New()
	}
	_, err = io.Copy(hash, reader)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s: %x", *algo, hash.Sum(nil))
}
