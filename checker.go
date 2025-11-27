package main

import (
	"net/http"
	"strings"
	"time"
)

func CheckLink(link string) string {
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "http://" + link
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(link)
	if err != nil {
		return "not available"
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return "available"
	}

	return "not available"
}
