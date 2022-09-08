package main

type Post struct {
	Draft      State
	Moderation State
	Published  State

	Text string

	CurrState State
}

func newPost(text string) *Post {
	p := &Post{
		Text: text,
	}
	draftState := &DraftState{post: p}
	moderationState := &ModerationState{post: p}
	publishedState := &PublishedState{post: p}
	p.setState(draftState)
	p.Draft = draftState
	p.Moderation = moderationState
	p.Published = publishedState
	return p
}

func (p *Post) setState(s State) {
	p.CurrState = s
}

func (p *Post) publish() error {
	return p.CurrState.publish()
}
