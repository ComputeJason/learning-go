package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/jasonchen/blogposts"
)

// A VERY GOOD CHAPTER again! read again.

// users looking at the code below can easily understand how to use our code and
// can simply do the following below

// func main() {
// 	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(posts)
// }

// notice how similar the lines are:
// main: posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
// test: posts, err := blogposts.NewPostsFromFS(fs)

// because of dependency injection, we can easily test and demonstrate functionality
// without directly depending on external dependencies.

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
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
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})

}