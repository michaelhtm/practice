package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func main() {

	client := github.NewClient(nil)

	newPR := &github.NewPullRequest{
		Title: github.String("This is a cool PR"),
		Head:  github.String("something-new"),
		Base:  github.String("origin"),
	}

	pr,resp, err := client.PullRequests.Create(context.Background(), "michaelhtm", "practice", newPR)
	if err != nil {
		fmt.Println(resp)
		fmt.Printf("yo stop %s\n", err)
		return
	}

	fmt.Printf("PR created: %s\n", pr.GetHTMLURL())
}
