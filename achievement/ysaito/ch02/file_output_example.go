package main

import "os"

func main() {
	file, err := os.Create("../tmp/test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File example\n"))
	file.Close()
}
