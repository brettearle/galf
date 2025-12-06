package testutil

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	ColorBlue    = "\033[34m"
	ColorGold    = "\033[33m" // gold/yellow
	ColorReset   = "\033[0m"
	ColorRedBold = "\033[1;31m"
)

func Blue(msg string) string {
	return fmt.Sprintf("%s%s%s", ColorBlue, msg, ColorReset)
}

func Gold(msg string) string {
	return fmt.Sprintf("%s%s%s", ColorGold, msg, ColorReset)
}

func Red(msg string) string {
	return fmt.Sprintf("%s%s%s", ColorRedBold, msg, ColorReset)
}

func WaitForReady(
	ctx context.Context,
	timeout time.Duration,
	endpoint string,
) error {
	client := http.Client{}
	start := time.Now()

	for {
		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodGet,
			endpoint,
			nil,
		)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			continue
		}

		if resp.StatusCode == http.StatusOK {
			fmt.Println("Endpoint is ready!")
			resp.Body.Close()
			return nil
		}
		resp.Body.Close()

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if time.Since(start) >= timeout {
				return fmt.Errorf("timeout reached while waiting for endpoint")
			}
			time.Sleep(250 * time.Millisecond)
		}
	}
}
