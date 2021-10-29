package main

import (
	"github.com/turnage/graw/reddit"
)

func NewNotifier(rctx *RedditContext, fn NotifyOutput) *Notifier {
	return &Notifier{
		Ctx:        rctx,
		NotifyFunc: fn,
	}
}

func (n *Notifier) Post(p *reddit.Post) error {
	msg := RedditMessage{
		Title:   p.Title,
		Content: p.SelfText,
		URL:     p.URL,
		Tag:     "Post",
	}

	n.NotifyFunc(msg)

	return nil
}

func (n *Notifier) UserPost(p *reddit.Post) error {
	msg := RedditMessage{
		Title:   p.Title,
		Content: p.SelfText,
		URL:     p.URL,
		Tag:     "Post",
	}

	n.NotifyFunc(msg)

	return nil
}

func (n *Notifier) UserComment(c *reddit.Comment) error {
	msg := RedditMessage{
		Title:   c.LinkTitle,
		Content: c.Body,
		URL:     c.LinkURL,
		Tag:     "Comment",
	}

	n.NotifyFunc(msg)

	return nil
}
