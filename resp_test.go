package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestResp(t *testing.T) {
	input := "$5\r\ngavin\r\n"
	reader := bufio.NewReader(strings.NewReader(input))

	b, _ := reader.ReadByte()

	if b != '$' {
		fmt.Println("Invalid type, expecting bulk strings only")
		os.Exit(1)
	}

	size, _ := reader.ReadByte()

	strSize, _ := strconv.ParseInt(string(size), 10, 64)

	reader.ReadByte()
	reader.ReadByte()

	message := make([]byte, strSize)
	reader.Read(message)

	fmt.Println(string(message))

}
