package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

func main() {
	rctx, err := ParseConfig("config.yml")
	if err != nil {
		return
	}

	apiHandle, err := reddit.NewScript(rctx.UserAgent, 5*time.Second)
	if err != nil {
		fmt.Printf("Script initialization error: %v\n", err)
		return
	}

	notifier := NewNotifier(rctx, PrintTerm, apiHandle)

	stop, wait, err := graw.Scan(notifier, apiHandle, rctx.GrawCfg)
	if err != nil {
		fmt.Printf("Scanner error: %v\n", err)
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		stop()
		fmt.Println("Sigint caught, exiting")
		os.Exit(1)
	}()

	if err := wait(); err != nil {
		fmt.Printf("Scanner error: %v\n", err)
	}
}

func PrintTerm(msg RedditMessage) {
	fmt.Printf("(%s) [%s] %s in /r/%s\n/u/%s\nPermalink: https://reddit.com%s\nContent URL: %s\n%s\n---\n", msg.Created.Format("01-02-2006 15:04:05"), msg.Tag, msg.Title, msg.Subreddit, msg.User, msg.Permalink, msg.ContentURL, msg.Content)
}
