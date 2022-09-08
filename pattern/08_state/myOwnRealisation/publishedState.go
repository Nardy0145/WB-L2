package main

import "fmt"

type PublishedState struct {
	post *Post
}

func (m *PublishedState) publish() error {
	fmt.Println("Post is already published!")
	m.post.setState(m.post.Draft)
	return nil
}
