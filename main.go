package main

import "fmt"

type Reader interface {
	read()
}

type Writer interface {
	write(string)
}

type ReaderWriter interface {
	Reader
	Writer
}

func writeToStream(writer ReaderWriter, text string) {
	writer.write(text)
}

func ReadFromStream(reader ReaderWriter) {
	reader.read()
}

// структура файл
type File struct {
	text string
}

// реализация методов типа *File
func (f *File) read() {
	fmt.Println(f.text)
}

func (f *File) write(message string) {
	f.text = message
	fmt.Println("Запись в файл строки", message)
}

func main() {
	myFile := &File{}
	writeToStream(myFile, "hello world")
	ReadFromStream(myFile)

	writeToStream(myFile, "lolly bomb")
	ReadFromStream(myFile)
}
