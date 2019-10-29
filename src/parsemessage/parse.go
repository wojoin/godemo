package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type alphaReader struct {
	reader io.Reader
}

func newAlphaReader(reader io.Reader) *alphaReader {
	return &alphaReader{reader: reader}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}

	copy(p, buf)
	return n, nil
}

type Header struct {
	number    int32
	separator byte
	//	UserDataMaxSize uint32
	//	HeaderOffset    uint32
	//	UserDataSize    uint32
	//	_               [5]byte
	//	Starcraft2      [22]byte
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

func main() {

	content, err := ioutil.ReadFile("number_0.data")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s\n\n", content)
	n := 0
	//length := int((len(content) - 4) / 8)
	length := len(content)
	fmt.Println("length: ", length)

	for n != length {
		fmt.Println("value: ", content[n:n+4])
		n = n + 5
	}

	fmt.Printf("File contents: %s\n\n", content[0:4])

	file, err := os.Open("number_0.data")
	var offset int64

	size, err := file.Seek(0, os.SEEK_END)
	fmt.Println("size: ", size)
	contentdata := make([]byte, size)

	n, er := file.Read(contentdata)
	if er != nil {
		fmt.Println(er.Error())
	}
	fmt.Println("read size: ", n)

	header := Header{}
	data := readNextBytes(file, 5) // 4 * uint32 (3) + 5 * byte (1) + 22 * byte (1) = 43

	buffer := bytes.NewBuffer(data)
	err = binary.Read(buffer, binary.LittleEndian, &header)

	if err != nil {
		log.Fatal("binary.Read failed", err)
	}
	//offset = offset + 8
	file.Seek(offset, 0)

	fmt.Printf("Parsed data:\n%+v\n", header)

}

//func main() {
//	// use an os.File as source for alphaReader
//	file, err := os.Open("./number_0.data")
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	defer file.Close()

//	reader := newAlphaReader(file)
//	p := make([]byte, 4)
//	for {
//		n, err := reader.Read(p)
//		if err == io.EOF {
//			break
//		}
//		fmt.Print(string(p[:n]))
//	}
//	fmt.Println()
//}
