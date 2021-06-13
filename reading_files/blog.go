package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)


type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

// Use the file system interface from the stdlib
// This interface has the Open method
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, file := range dir {
		post, err := getPost(fileSystem, file.Name())

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil

}

// Seperating the 'opening file code' from 'parsing file contents code' will make it simpler to work with
// Single Responseibility Principle
func getPost(fileSystem fs.FS, filename string) (Post, error) {
	postFile, err := fileSystem.Open(filename)

	if err != nil {
		return Post{}, nil
	}

	defer postFile.Close()
	return newPost(postFile)
}

// When refactoring out new funcs/methods, think about the arguments.
// Think about coupling and conhesion.
// Does newPost have to be coupled to an fs.File as the file type?
// Since we're not using a lot of the methods on fs.File type,
// we should loosen the coupling and use the io.Reader interface

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
	tagSeperator         = "Tags: "
)

// Theres a readable way of trimming a prefix from a string with the func
// strings.TrimPrefix

func newPost(postFile io.Reader) (Post, error) {
	// The stdlib has a nice lib to scan through data line by line; bufio.Scanner
	// It also takes an io.Reader to read
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		// This returns the string without the prefix
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	titleLine := readMetaLine(titleSeperator)
	descriptionLine := readMetaLine(descriptionSeperator)
	tagLine := readMetaLine("Tags: ")
	tags := strings.Split(tagLine, ", ")

	body := readBody(scanner)

	post := Post{Title: titleLine, Description: descriptionLine, Tags: tags, Body: body}

	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	buf := bytes.Buffer{}

	// Scan() returns a bool which indicates whether there's more data to scan so we can
	// use this in a loop to scan the data until the end.
	// After every scan, we write data into the buffer using Fprintln
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	// Trim trailing newline at end of buffer
	return strings.TrimSuffix(buf.String(), "\n")
}
