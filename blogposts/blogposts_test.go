package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/jdra000/learn-go-with-tests/blogposts"
)

func TestNewBlogpost(t *testing.T){
    const ( firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello World`
        secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
    )

    // Package fstest implements support for testing implementations ans users of file systems
    fs := fstest.MapFS {
        "file-1.md": {Data: []byte(firstBody)},
        "file-2.md": {Data: []byte(secondBody)},
    }

    posts, err := blogposts.NewPostFromFS(fs)

    if err != nil {
        t.Fatal(err)
    }

    if len(posts) != len(fs){
        t.Errorf("got %d posts wanted %d posts", len(posts), len(fs))
    }

    assertPost(t, posts[0], blogposts.Post{
        Title:          "Post 1",
        Description:    "Description 1",
        Tags:           []string{"tdd", "go"},
        Body:           "Hello World",
    })
}
func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
   t.Helper()
   if !reflect.DeepEqual(got, want) {
        t.Errorf("got %+v want %+v", got, want)
   }
}
