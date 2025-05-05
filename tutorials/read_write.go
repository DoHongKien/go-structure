package main

import (
	"os"
)

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(filePath string) {
	data, err := os.ReadFile(filePath)
	checkPanic(err)
	println("File content:", string(data))
}

func WriteFile(filePath string, data []byte) {
	err := os.WriteFile(filePath, data, 0644)
	checkPanic(err)
	println("Data written to file:", filePath)
}

func main() {
	fileName := "output.txt"
	WriteFile(fileName, []byte("Hello, World!"))
	ReadFile(fileName)
}
