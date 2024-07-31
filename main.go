package main

import (
	"flag"
	"fmt"
	"golang-interface-modular/backend/file"
	"golang-interface-modular/backend/s3"
	"os"
)

func ProcessData(backend Backend) error {
	backend.Save("test", []byte("Here is a string...."))
	return nil
}

func main() {
	storageFlag := flag.String("backend", "file", "Storage backend to use")

	flag.Parse()
	if *storageFlag == "" {
		flag.Usage()
		os.Exit(1)
	}

	var backend Backend

	switch *storageFlag {
	case "file":
		backend = file.Storage{}
	case "s3":
		backend = s3.Storage{}
	default:
		fmt.Println("unknown storage backend specified")
		os.Exit(1)
	}

	ProcessData(backend)
}
