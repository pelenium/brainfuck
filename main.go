package main

import (
	"os"
)

func main() {
	filePath := os.Args[1]

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	src := string(data)

	fuckBrain(src)
}
