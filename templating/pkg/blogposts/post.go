package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func (p *Post) Route() string {
	var tmp string = p.Title
	tmp = strings.ReplaceAll(tmp, " ", "-")
	tmp = strings.ToLower(tmp)
	tmp = strings.Trim(tmp, "-")
	return "/post/" + tmp
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	title := func() string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), "Title: ")
	}()

	description := func() string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), "Description: ")
	}()

	tags := func() []string {
		scanner.Scan()
		str := strings.TrimPrefix(scanner.Text(), "Tags: ")
		return strings.Split(str, ", ")
	}()

	body := func() string {
		scanner.Scan() // skip line
		buf := bytes.Buffer{}
		for scanner.Scan() {
			fmt.Fprintln(&buf, scanner.Text())
		}
		return strings.TrimSuffix(buf.String(), "\n")
	}()

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}, nil
}
