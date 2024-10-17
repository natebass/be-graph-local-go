package gh

import (
	"context"
	"fmt"

	"github.com/google/go-github/v64/github"
)

func WillNorris() {
	client := github.NewClient(nil)

	// list all organizations for user "willnorris"
	orgs, _, err := client.Organizations.List(context.Background(), "willnorris", nil)
	fmt.Println(orgs)
	fmt.Println(err)
}
