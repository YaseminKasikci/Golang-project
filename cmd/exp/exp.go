package main

import (
	"context"
	"fmt"
	"strings"
)

type ctxKey string

const (
	favoriteColorKey ctxKey = "favorite-color"
)

func main() {
	ctx := context.Background()
	// set color to blue
	ctx = context.WithValue(ctx, favoriteColorKey, "blue")

	value := ctx.Value(favoriteColorKey)

	intValue, ok := value.(int)
	if !ok {
		fmt.Println("it isn't an int")
	} else {
		fmt.Println(intValue + 4)
	}

	strValue, ok := value.(string)
	if !ok {
		fmt.Println("it isn't a string")
	} else {
		fmt.Println(strings.HasPrefix(strValue, "b"))
	}

	// fmt.Println(strValue)s
	// fmt.Println(strings.HasPrefix(strValue, "b"))
}
