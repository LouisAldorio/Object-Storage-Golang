package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

func PathToFile(fullPath string, fileType string) ([]*graphql.Upload, error) {
	var filePack []*graphql.Upload

	filename := strings.Split(fullPath, "/")[len(strings.Split(fullPath, "/"))-1]
	fmt.Println(filename)

	fileName := fullPath
	pdfFileOpen, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	readerIoUtil, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	fileStatus, err := pdfFileOpen.Stat()

	if err != nil {
		return nil, err
	}

	var r io.Reader
	r = bytes.NewReader(readerIoUtil)

	fileSize := fileStatus.Size()

	fmt.Println(fileName, "asd")

	tempFile := graphql.Upload{
		File:        r,
		Filename:    filename,
		Size:        fileSize,
		ContentType: fileType,
	}

	filePack = append(filePack, &tempFile)

	fmt.Println(filePack[0].Filename)

	return filePack, nil
}
