package main

import "net/url"

type Config struct {
	Address url.URL
	Output  Output
	Random  bool
}
