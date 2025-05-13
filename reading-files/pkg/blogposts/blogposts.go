package blogposts

import "io/fs"

func NewPostsFromFilesystem(filesystem fs.FS) ([]Post, error) {
	cwd, err := fs.ReadDir(filesystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for range cwd {
		posts = append(posts, Post{})
	}
	return posts, nil
}
