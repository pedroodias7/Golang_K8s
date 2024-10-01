package main

import (
	"io"
	"fmt"
	"log"
)


//json to Struct Golang
type MySlowReader struct {
	content string
	pos int
}


func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.pos+1 <= len(m.content) {
		n := copy(p, m.content[m.pos:+1])
		m.pos++
		return n, nil 
	}
	return 0, io.EOF
}

func main() {

	mySlowReaderInstance := &MySlowReader{
		Content: "Hello word",
	}


	out, err := io.ReadAll(mySlowReaderInstance)

	if err != nil {
		log.Fatal(err)
	}
	

	fmt.Printf("output: %s\n", out)



}
