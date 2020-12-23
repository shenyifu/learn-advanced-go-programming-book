package main

import "context"
import "fmt"

func GenerateNature(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				return
			}

		}
	}()
	return ch
}

func PrimFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)

	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case <-ctx.Done():
					return
				case out <- i:
				}
			}
		}
	}()
	return out
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := GenerateNature(ctx)

	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v : %v\n", i+1, prime)
		ch = PrimFilter(ctx, ch, prime)
	}
	cancel()
}
