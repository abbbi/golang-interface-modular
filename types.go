package main

type Backend interface {
	Save(filename string, data []byte) error
}
