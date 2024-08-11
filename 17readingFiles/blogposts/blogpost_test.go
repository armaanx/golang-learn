package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/armaanx/blogposts"
)

type StubFailingFs struct {
}

func (s StubFailingFs) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i fail")
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)
	fs := fstest.MapFS{
		"hello-world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFs(fs)

	if err != nil {
		t.Fatal(err)
	}

	got := posts[0]
	want := blogposts.Post{Title: "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	}

	assertPost(t, got, want)

	if len(posts) != len(fs) {

		t.Errorf("got %d posts; wanted %d posts", len(posts), len(fs))
	}
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
