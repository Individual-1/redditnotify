package main

import (
	"strings"
	"time"

	"github.com/turnage/graw/reddit"
)

func NewNotifier(rctx *RedditContext, fn NotifyOutput, script reddit.Script) *Notifier {
	return &Notifier{
		Ctx:        rctx,
		NotifyFunc: fn,
		script:     script,
	}
}

func (n *Notifier) Post(p *reddit.Post) error {
	msg := RedditMessage{
		Title:      p.Title,
		Subreddit:  p.Subreddit,
		Permalink:  p.Permalink,
		User:       p.Author,
		Content:    p.SelfText,
		ContentURL: p.URL,
		Tag:        "Post",
		Created:    time.Unix(int64(p.CreatedUTC), 0),
	}

	n.NotifyFunc(msg)

	return nil
}

func (n *Notifier) UserPost(p *reddit.Post) error {
	uctx, ok := n.Ctx.Users[strings.ToLower(p.Author)]
	if !ok {
		//return errors.New("User comment received for unregistered user")
		return nil
	}

	_, ok = uctx.Subreddits[strings.ToLower(p.Subreddit)]
	if uctx.IsAllow != ok {
		// Irrelevant subreddit, silently drop
		return nil
	}

	msg := RedditMessage{
		Title:      p.Title,
		Subreddit:  p.Subreddit,
		Permalink:  p.Permalink,
		User:       p.Author,
		Content:    p.SelfText,
		ContentURL: p.URL,
		Tag:        "Post",
		Created:    time.Unix(int64(p.CreatedUTC), 0),
	}

	n.NotifyFunc(msg)

	return nil
}

func (n *Notifier) UserComment(c *reddit.Comment) error {
	uctx, ok := n.Ctx.Users[strings.ToLower(c.Author)]
	if !ok {
		//return errors.New("User comment received for unregistered user")
		return nil
	}

	_, ok = uctx.Subreddits[strings.ToLower(c.Subreddit)]
	if uctx.IsAllow != ok {
		// Irrelevant subreddit, silently drop
		return nil
	}

	msg := RedditMessage{
		Title:      c.LinkTitle,
		Subreddit:  c.Subreddit,
		Permalink:  c.Permalink,
		User:       c.Author,
		Content:    c.Body,
		ContentURL: c.LinkURL,
		Tag:        "Comment",
		Created:    time.Unix(int64(c.CreatedUTC), 0),
	}

	n.NotifyFunc(msg)

	return nil
}
