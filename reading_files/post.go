package readingfiles

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	tittle := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagsSeparator), ", ")

	return Post{
		Title:       tittle,
		Description: description,
		Tags:        tags,
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
