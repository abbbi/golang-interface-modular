package s3

import "fmt"

type Storage struct{}

func (s3 Storage) Save(filename string, data []byte) error {
	fmt.Printf("Saving data to s3 object: %s\n", filename)
	return nil
}
