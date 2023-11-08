package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("myfile.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufferedWriteer := bufio.NewWriter(file)
	bs := []byte("abcd efg hijk lmnop qrsn tuv wxy yanz")
	bytesWritten, err := bufferedWriteer.Write(bs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file) %d \n", bytesWritten)
	bytesAvailable := bufferedWriteer.Available()
	log.Println("Bytes available in buffer: ", bytesAvailable)

	bytesWritten, err = bufferedWriteer.WriteString("\njust a random string")
	if err != nil {
		log.Fatal(err)
	}
	unflushedBufferSize := bufferedWriteer.Buffered()
	log.Printf("Bytes buffered : %d \n", unflushedBufferSize)
	bufferedWriteer.Flush()

}
