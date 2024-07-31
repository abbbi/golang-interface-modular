package main

import (
	"flag"
	"fmt"
	"golang-interface-modular/backend/file"
	"golang-interface-modular/backend/s3"
	"os"
)

func ProcessStorage(s Storage) error {
	s.Save("test", []byte("Here is a string...."))
	return nil
}

func main() {
	storageFlag := flag.String("backend", "file", "Storage backend to use: default: file")

	flag.Parse()
	if *storageFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	var backend Storage

	switch *storageFlag {
	case "file":
		backend = file.Storage{}
	case "s3":
		backend = s3.Storage{}
	default:
		fmt.Println("unknown storage backend specified")
		os.Exit(1)
	}

	ProcessStorage(backend)
}
