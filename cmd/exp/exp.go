package main

import (
	stdctx "context"
	"fmt"

	"github/yaseminkasikci/lenslocked/context"
	"github/yaseminkasikci/lenslocked/models"
)

type ctxKey string

const (
	favoriteColorKey ctxKey = "favorite-color"
)

func main() {
	ctx := stdctx.Background()
	// set color to blue
	user := models.User{
		Email: "bobo@gmail.com",
	}
	ctx = context.WithUser(ctx, &user)

	retrievedUser := context.User(ctx)
	fmt.Println(retrievedUser.Email)
}
