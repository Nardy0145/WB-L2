package main

import "fmt"

type ModerationState struct {
	post *Post
}

func (m *ModerationState) publish() error {
	fmt.Printf("Successfully published post: %s\n", m.post.Text)
	m.post.setState(m.post.Published)
	return nil
}
