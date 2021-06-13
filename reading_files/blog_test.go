package blog

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

// Requirements:
// Read in a folder of blog_posts.md files, convert blog posts to struct and server via a webserver
// In go, backticks can span multiple lines and ignore special characters (e.g. \n)
func TestNewBlogPosts(t *testing.T) {
	// fs = filesystem
	// MapFS is a simple in-memory file system to use in tests, represented as a map (path name; file data)
	// This is better than maintaining a list of test files

	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: python
---
Hello
Mike`
	)

	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts but wanted %d posts", len(posts), len(fs))
	}

	got := posts[0]

	wantBody := `Hello
World`

	want := Post{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: wantBody}

	assertPost(t, got, want)

}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

}

// We can use this interface/method to test errors when reading files
type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("I always fail when called")
}
