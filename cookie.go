package knot

import (
	"net/http"
	"net/url"
	"time"
)

func (r *Request) Cookie(name string) (*http.Cookie, error) {
	c, e := r.HttpRequest().Cookie(name)
	if e == nil {
		r.ResponseConfig().Cookies()[name] = c
	}
	return c, nil
}

func (r *Request) SetCookie(name string, value string, expiresAfter time.Duration) {
	c := &http.Cookie{}
	c.Name = name
	c.Value = value
	u, e := url.Parse(r.HttpRequest().URL.String())
	if e == nil {
		c.Expires = time.Now().Add(expiresAfter)
		c.Domain = u.Host
		r.ResponseConfig().Cookies()[name] = c
	}
}
