package main

import "fmt"

type DraftState struct {
	post *Post
}

func (d *DraftState) publish() error {
	fmt.Println("Sent publication to moderation!")
	d.post.setState(d.post.Moderation)
	return nil
}
