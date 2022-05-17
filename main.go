package main

import "fmt"

type Reader interface {
	read()
}

type Writer interface {
	write(string)
}

func writeToStream(writer Writer, text string) {
	writer.write(text)
}

func ReadFromStream(reader Reader) {
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
}
