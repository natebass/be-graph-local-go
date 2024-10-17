package gh

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v64/github"
	"golang.org/x/term"
)

func auth() {
	fmt.Print("GitHub Token: ")
	byteToken, _ := term.ReadPassword(int(os.Stdin.Fd()))
	println()
	token := string(byteToken)

	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)

	user, resp, err := client.Users.Get(ctx, "")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	// Rate.Limit should most likely be 5000 when authorized.
	log.Printf("Rate: %#v\n", resp.Rate)

	// If a Token Expiration has been set, it will be displayed.
	if !resp.TokenExpiration.IsZero() {
		log.Printf("Token Expiration: %v\n", resp.TokenExpiration)
	}

	fmt.Printf("\n%v\n", github.Stringify(user))
}
