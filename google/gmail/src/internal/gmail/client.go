package gmail

import (
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"time"

	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/googleapi"
)

type Client struct {
	srv *gmail.Service
}

func NewClient(srv *gmail.Service) *Client {
	return &Client{srv: srv}
}

func withRetries[T any](operation func() (T, error)) (T, error) {
	var maxRetries = 6
	var baseDelay = 1 * time.Second

	for i := 0; i < maxRetries; i++ {
		res, err := operation()
		if err == nil {
			return res, nil
		}
		
		if apiErr, ok := err.(*googleapi.Error); ok && (apiErr.Code == 403 || apiErr.Code == 429 || apiErr.Code >= 500) {
			delay := float64(baseDelay) * math.Pow(2, float64(i))
			jitter := (rand.Float64() * 0.2) + 0.9 // 90% to 110%
			time.Sleep(time.Duration(delay * jitter))
			continue
		}
		var empty T
		return empty, err
	}
	var empty T
	return empty, fmt.Errorf("max retries exceeded")
}

// StreamMessages fetches messages with pagination and pushes to a channel.
func (c *Client) StreamMessages(query string, out chan<- *gmail.Message, errs chan<- error) {
	defer close(out)
	defer close(errs)

	pageToken := ""
	for {
		req := c.srv.Users.Messages.List("me").Q(query).MaxResults(500)
		if pageToken != "" {
			req.PageToken(pageToken)
		}
		
		res, err := withRetries(func() (*gmail.ListMessagesResponse, error) {
			return req.Do()
		})
		
		if err != nil {
			errs <- fmt.Errorf("list messages failed: %w", err)
			return
		}

		for _, m := range res.Messages {
			out <- m
		}

		if res.NextPageToken == "" {
			break
		}
		pageToken = res.NextPageToken
	}
}

// GetRawMessageAndDate fetches the raw .eml content and the native Gmail InternalDate accurately.
func (c *Client) GetRawMessageAndDate(id string) ([]byte, time.Time, error) {
	msg, err := withRetries(func() (*gmail.Message, error) {
		return c.srv.Users.Messages.Get("me", id).Format("raw").Do()
	})
	if err != nil {
		return nil, time.Time{}, err
	}
	
	decoded, err := base64.URLEncoding.DecodeString(msg.Raw)
	if err != nil {
		return nil, time.Time{}, fmt.Errorf("decode error %s: %w", id, err)
	}
	
	sentAt := time.UnixMilli(msg.InternalDate)
	return decoded, sentAt, nil
}
