package main

import (
	"net/url"
	"time"
)

type Config struct {
	Address   url.URL
	Output    Output
	Random    bool
	ExpiresIn time.Duration
}
