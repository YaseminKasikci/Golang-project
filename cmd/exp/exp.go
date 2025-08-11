package main

import (
	"context"
	"fmt"
)

type ctxKey string

const (
	favoriteColorKey ctxKey = "favorite-color"
)

func main() {
	ctx := context.Background()
	// set color to blue
	ctx = context.WithValue(ctx, favoriteColorKey, "blue")

	// some other package writes a favorite-color
	ctx = context.WithValue(ctx, "favorite-color", 132)

	//oh not!
	value1 := ctx.Value(favoriteColorKey)
	value2 := ctx.Value("favorite-color")
	fmt.Println(value1, value2)
}
