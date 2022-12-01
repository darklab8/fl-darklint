/*
# File handling functions

F in OpenToReadF stands for... Do succesfully, or log to Fatal level and exit
*/
package utils

import (
	"bufio"
	"os"
)

type File struct {
	Filepath string
	file     *os.File
}

func (f File) OpenToReadF() File {
	file, err := os.Open(f.Filepath)
	f.file = file

	CheckFatal(err, "failed to open ", f.Filepath)
	return f
}

func (f File) Close() {
	f.file.Close()
}

func (f File) ReadLines() []string {

	scanner := bufio.NewScanner(f.file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func (f File) CreateToWriteF() File {
	file, err := os.Create(f.Filepath)
	f.file = file
	CheckFatal(err, "failed to open ", f.Filepath)

	return f
}
func (f File) WriteF(msg string) {
	_, err := f.file.WriteString(msg)

	CheckFatal(err, "failed to write string to file")
}
