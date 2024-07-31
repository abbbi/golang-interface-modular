package file

import "fmt"

type Storage struct{}

func (file Storage) Save(filename string, data []byte) error {
	fmt.Printf("Saving data to file: %s\n", filename)
	return nil
}
