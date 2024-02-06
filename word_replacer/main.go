package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	input, err := os.ReadFile("original.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output := bytes.Replace(input, []byte("However"), []byte("Twitter"), -1)

	if err = os.WriteFile("newfile.txt", output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
