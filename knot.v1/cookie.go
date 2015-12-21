package knot

import (
	"net/http"
	"net/url"
	"time"
)

var DefaultCookieExpire time.Duration

func (r *WebContext) Cookie(name string, def string) (*http.Cookie, bool) {
	exist := true
	c, e := r.Request.Cookie(name)
	if e == nil && def != "" {
		if int(DefaultCookieExpire) == 0 {
			DefaultCookieExpire = 30 * 24 * time.Hour
		}
		r.SetCookie(name, def, DefaultCookieExpire)
		exist = false
	}
	return c, exist
}

func (r *WebContext) SetCookie(name string, value string, expiresAfter time.Duration) {
	c := &http.Cookie{}
	c.Name = name
	c.Value = value
	u, e := url.Parse(r.Request.URL.String())
	if e == nil {
		c.Expires = time.Now().Add(expiresAfter)
		c.Domain = u.Host
	}
	if r.cookies == nil {
		r.cookies = map[string]*http.Cookie{}
	}
	r.cookies[name] = c
}

func (r *WebContext) Cookies() map[string]*http.Cookie {
	if r.cookies == nil {
		r.cookies = map[string]*http.Cookie{}
	}
	return r.cookies
}
