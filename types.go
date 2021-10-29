package main

import (
	"github.com/turnage/graw"
)

type Notifier struct {
	Ctx        *RedditContext
	NotifyFunc NotifyOutput
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
	GrawCfg graw.Config     // Graw config
	Users   map[string]User // User details for filtering
}

type User struct {
	Name       string              // User name
	IsAllow    bool                // Whether subreddits is a Denylist or Allowlist
	Subreddits map[string]struct{} // List of subs to include/exclude
}

type RedditMessage struct {
	Title   string
	Content string
	URL     string
	Tag     string
}
