// Package main for HTTP client example
package main

import (
	"log"

	"github.com/ShiftOver/shiftover-backend/pkg/httpclient"
)

func main() {
	c := httpclient.NewClient(httpclient.SetMaxRetry(3), httpclient.SetTimeout(10))

	response, err := c.Get("http://example.com", httpclient.SendOptions{
		"queries": map[string]string{
			"queries": "example",
			"limit":   "10",
		},
		"headers": map[string]string{
			"custom":       "example",
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		log.Panicf("error - [main]: unable to get response: %v", err)
	}
	_ = response
}
