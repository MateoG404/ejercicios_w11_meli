package main

import (
	"context"
	"fmt"
)

type contextKey string

func main() {
	// Context creationff
	ctx := context.Background()

	// Add value to context
	ctx = context.WithValue(ctx, contextKey("key"), "value")
	ctx = context.WithValue(ctx, contextKey("api-key"), "super-secret-key")

	// Get value from context
	value := ctx.Value(contextKey("key"))
	fmt.Println(value)

}
