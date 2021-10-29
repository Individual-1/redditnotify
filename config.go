package main

import (
	"os"

	"github.com/turnage/graw"
	"gopkg.in/yaml.v3"
)

func ParseConfig(filename string) (*RedditContext, error) {
	config := Config{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	rctx := GenerateGrawConfig(config)
	return rctx, nil
}

func GenerateGrawConfig(config Config) *RedditContext {
	rctx := RedditContext{}
	rctx.GrawCfg = graw.Config{}
	rctx.Users = make(map[string]User)

	rctx.UserAgent = config.UserAgent

	rctx.GrawCfg.Subreddits = config.Targets.Subreddits
	rctx.GrawCfg.Users = make([]string, 0, 10)

	for _, u := range config.Targets.Users {
		user := User{}
		user.Name = u.Name
		user.IsAllow = u.IsAllow

		user.Subreddits = make(map[string]struct{})
		for _, s := range u.Subreddits {
			user.Subreddits[s] = struct{}{}
		}

		rctx.Users[u.Name] = user
		rctx.GrawCfg.Users = append(rctx.GrawCfg.Users, u.Name)
	}

	return &rctx
}

/*
Non-stream-based model

type RedditContext struct {
	Subreddits          map[string]Subreddit
	Users               map[string]User
	LastHomeId          string // Last home screen id seen
	LastFriendPostId    string // Last friend post id seen
	LastFriendCommentId string // Last friend comment id seen
}

type Subreddit struct {
	Name   string // Name of subreddit
	LastId string // Last id retrieved
}

type User struct {
	Name          string              // User name
	IsAllow       bool                // Whether subreddits is a Denylist or Allowlist
	Subreddits    map[string]struct{} // List of subs to include/exclude
	LastPostId    string              // Last post id seen
	LastCommentId string              // Last comment id seen
}

func NewRedditContext() *RedditContext {
	rc := RedditContext{}
	rc.Subreddits = make(map[string]Subreddit)
	rc.Users = make(map[string]User)

	return &rc
}

func ParseConfig(filename string) error {
	config := Config{}

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}

}

func GenerateRedditContext(config Config, client ) (*RedditContext, error) {
	rc := NewRedditContext()

	for _, sub := range config.Targets.Subreddits {
		s := Subreddit{
			Name:   sub,
			LastId: "",
		}

		rc.Subreddits[sub] = s
	}

	for _, user := range config.Targets.Users {
		subs := make(map[string]struct{})

		for _, sub := range user.Subreddits {
			subs[sub] = struct{}{}
		}

		u := User{
			Name:          user.Name,
			IsAllow:       user.IsAllow,
			Subreddits:    subs,
			LastPostId:    "",
			LastCommentId: "",
		}

		rc.Users[user.Name] = u
	}

	if client != nil {
		friends, _, err := client.Account.Friends(context.Background())
		if err != nil {
			return nil, err
		}

		friendMap := make(map[string]struct{})
		for _, r := range relations {
			friendMap[r.User] = struct{}{}
		}

		// Add all users in our target list to friends
		for name, _ := range rc.Users {
			if _, present := friendMap[name]; !present {
				_, _, err = client.User.Friend(context.Background(), name)
				if err != nil {
					return nil, err
				}
			}
		}

		subs, _, err := client.
		subMap := make(map[string]struct{})


		// Make sure we are subscribed to all target subs
		for sub, _ := range rc.Subreddits {

		}

	}

	return rc, nil
}
*/
