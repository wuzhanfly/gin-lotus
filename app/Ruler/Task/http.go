package Task

import (
	"context"
	"net/http"
	"time"
)

type fn func(ctx context.Context) result

type result struct {
	Val interface{}
	Err error
}

func doWithTimeout(ctx context.Context, fn fn) result {
	ch := make(chan result)
	go func(ctx context.Context, ch chan<- result) {
		ch <- fn(ctx)
	}(ctx, ch)

	select {
	case <-ctx.Done(): // timeout
		go func() { <-ch }() // wait ch return...
		return result{Err: ctx.Err()}
	case res := <-ch: // normal case
		return res
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	result := doWithTimeout(ctx, func(ctx context.Context) result {
		// replace with your own logic!
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://filfox.info/api/v1/miner/list/power?pageSize=25&page=0&continent=AS", nil)
		resp, err := http.DefaultClient.Do(req)
		return result{
			Val: resp,
			Err: err,
		}
	})
	switch {
	case ctx.Err() == context.DeadlineExceeded:
		// handle timeout
	case result.Err != nil:
		// handle logic error
	default:
		// do with result.Val
	}
}
