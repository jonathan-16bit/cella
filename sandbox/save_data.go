package main

import (
	"os"
	"fmt"
)

func SaveData1(path string, data []byte) error {
	// O_CREATE: create if needed, O_TRUNC: empty existing contents
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return err
	}

	// Close when function returns (including when write fails)
	defer fp.Close()

	// Write bytes, return error on fail
	_, err = fp.Write(data)
	if err != nil {
		return err
	}
	
	// fsync: asks OS to finish saving changes to storage
	return fp.Sync()
}

func main() {
	err := SaveData1("./test.txt", []byte("the quick brown fox jumps over the lazy dog"))
	if err != nil {
		fmt.Println(err)
	}
}
