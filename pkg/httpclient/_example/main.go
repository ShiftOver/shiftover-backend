// Package main for HTTP client example
package main

import (
	"time"

	"github.com/labstack/gommon/log"

	"github.com/ShiftOver/shiftover-backend/pkg/httpclient"
)

func main() {
	c := httpclient.NewClient(httpclient.SetMaxRetry(3), httpclient.SetTimeout(30 * time.Second))

	response, err := c.Get("https://66f0ebddf2a8bce81be6fcb0.mockapi.io/api/v1/test", httpclient.SendOptions{
		// "queries": map[string]string{
		// 	"queries": "example",
		// 	"limit":   "10",
		// },
		"headers": map[string]string{
			"custom":       "example",
			"Content-Type": "application/json",
		},
	})
	if err != nil {
		log.Panicf("error - [main]: unable to get response: %v", err)
	}

	log.Infof("Status: %v\n", response.StatusCode)
	log.Infof("Headers: %v\n", response.Header)
	log.Infof("Body: %v\n", string(response.Body))
}
