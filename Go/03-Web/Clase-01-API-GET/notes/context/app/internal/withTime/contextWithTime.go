package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func do_something(ctx context.Context) {
	fmt.Println("Doing something...")
	for i := 0; i < 3; i++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Doing something...")
		case <-ctx.Done():
			fmt.Println("Operation canceled by the time of the context")
			return
		}
	}
	fmt.Println("Operation finished")
}
func main() {
	// Creation context with time
	fmt.Println("Start")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Use context
	go do_something(ctx)

	// Wait for a key press to continue main
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	cancel()
	// Wait for a key press to continue main
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Println("End")
}
