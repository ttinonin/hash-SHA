package main

import (
	"fmt"
	"os"
	"crypto/sha256"
	"crypto/sha512"
	"strings"
)

func get_flags() []string {
	var flags []string

	for i := 1; i < len(os.Args); i++ {
		if strings.HasPrefix(os.Args[i], "-") {
			flags = append(flags, os.Args[i])	
		}
	}

	return flags
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: Missing arguments!")
		fmt.Fprintln(os.Stderr, "Usage: sha256 <string> [-384] [-512]")
		os.Exit(1)
	}

	word := os.Args[1]
	
	flags := get_flags()

	word_sha := sha256.Sum256([]byte(word))

	fmt.Printf("SHA256: %x\n", word_sha)

	for _, flag := range flags {
		switch flag {
		case "-384":
			fmt.Printf("SHA384: %x\n", sha512.Sum384([]byte(word)))	
		case "-512":
			fmt.Printf("SHA512: %x\n", sha512.Sum512([]byte(word)))	
		}
	}
}
