package main

import "fmt"

func main() {
	// byteArray は []byte{0x41, 0x53, 0x43, 0x49, 0x49}
	byteArray := []byte("ASCII")
	// str は "ASCII"
	str := string([]byte{0x41, 0x53, 0x43, 0x49, 0x49})

	fmt.Println(byteArray) // [65 83 67 73 73]
	fmt.Println(str)       // ASCII
}
