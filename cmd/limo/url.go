package main

import (
	"net/url"
)

const DefaultScheme = "https"

type URLFlag url.URL

func (f URLFlag) Type() string {
	return "string"
}

func (f *URLFlag) String() string {
	return (*url.URL)(f).String()
}

func (f *URLFlag) Set(s string) error {
	u, err := url.Parse(s)
	if err != nil {
		return err
	}

	if u.Host == "" {
		if u, err = url.Parse(DefaultScheme + "://" + s); err != nil {
			return err
		}
	}

	*f = (URLFlag)(*u)
	return nil
}
