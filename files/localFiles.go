package files

import (
	"fmt"
	"io"
	"os"
)

func LocalFiles() {
	fileName := "./files/fromString.txt"
	file, err := os.Create(fileName)
	defer file.Close()
	CheckError(err)
	length, err := io.WriteString(file, "Hello from Go!")
	fmt.Printf("Wrote a file with %v characters\n", length)
	readFile(fileName)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	CheckError(err)
	fmt.Println("Text read from file:", string(data))
}
