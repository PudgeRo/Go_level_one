package main

import (
	"fmt"
	"io"
	"log"
)

var (
	ErrWrite = fmt.Errorf("cannot write")
	ErrRead  = fmt.Errorf("cannot read")
)

type myType []byte

func (m myType) Read(p []byte) (int, error) {
	n := 0
	for i := 0; i < len(m); i++ {
		p[n] = m[i]
		n++
	}
	return n, io.EOF
}

func (m myType) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	for i := 0; i < len(p); i++ {
		m[i] = p[i]
	}
	return len(p), nil
}

func main() {
	greating := "Hi, it's me"
	m := make(myType, len(greating))
	_, err := io.WriteString(m, greating)
	if err != nil {
		log.Fatalln(ErrWrite)
	}
	newString, err := io.ReadAll(m)
	if err != nil {
		log.Fatalln(ErrRead)
	}
	fmt.Println(string(newString))
}
