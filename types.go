package main

import (
	"time"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type Notifier struct {
	Ctx        *RedditContext
	NotifyFunc NotifyOutput
	script     reddit.Script
}

type NotifyOutput func(RedditMessage)

type Config struct {
	Targets struct {
		Subreddits []string `yaml:"subreddits"`
		Users      []struct {
			Name       string   `yaml:"name"`
			IsAllow    bool     `yaml:"isAllow"`
			Subreddits []string `yaml:"subreddits"`
		} `yaml:"users"`
	} `yaml:"targets"`
	UserAgent string `yaml:"userAgent"`
}

type RedditContext struct {
	GrawCfg   graw.Config     // Graw config
	Users     map[string]User // User details for filtering
	UserAgent string          // User agent
}

type User struct {
	Name       string              // User name
	IsAllow    bool                // Whether subreddits is a Denylist or Allowlist
	Subreddits map[string]struct{} // List of subs to include/exclude
}

type RedditMessage struct {
	Title      string
	Subreddit  string
	Permalink  string
	User       string
	Content    string
	ContentURL string
	Tag        string
	Created    time.Time
}
