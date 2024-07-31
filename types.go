package main

type Storage interface {
	Save(filename string, data []byte) error
}
