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

	notifier := NewNotifier(rctx, PrintTerm)

	apiHandle, err := reddit.NewScript("test", 5*time.Second)
	if err != nil {
		fmt.Printf("Script initialization error: %v\n", err)
		return
	}

	stop, wait, err := graw.Scan(*notifier, apiHandle, rctx.GrawCfg)
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
	fmt.Printf("[%s] %s\n%s\n%s", msg.Tag, msg.Title, msg.URL, msg.Content)
}
